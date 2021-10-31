#!/bin/bash

#shellcheck disable=SC2034
test_name=":cypress: e2e tests"
# Note: we only run cypress here, ignore diagnostics and inspec
test_deploy_inspec_profiles=()
test_skip_diagnostics=true

do_create_config() {
    do_create_config_default
    #shellcheck disable=SC2154
    cat <<EOF >> "$test_config_path"
[global.v1.proxy]
host = "localhost"
port = 80
no_proxy = ["0.0.0.0", "127.0.0.1"]
EOF
}

# Note: these tests use the UI, so a valid license needs to be set up.
# Alternatively, the UI tests could also _apply the license_ and thus
# give us some test coverage for the license apply modal etc; but we
# are not there, yet.
do_deploy() {
    do_deploy_default
    log_info "applying dev license"
    chef-automate license apply "$A2_LICENSE"
    timestamp=$(date +"%m-%d-%y-%H-%M")

    log_info "generating admin token"
    if ! token=$(chef-automate iam token create "$timestamp-tok" --admin); then
        log_error "Non-zero exit code, output:"
        log_error "$token"
        return 1
    fi

    export CYPRESS_ADMIN_TOKEN=$token
    export CYPRESS_RUN_FLAKY=$FLAKY

    log_info "fixing dns resolution for '${CONTAINER_HOSTNAME}'"
    echo "127.0.0.1 ${CONTAINER_HOSTNAME}" >> /etc/hosts

    yum install openssh -y
    yum install openssh-server -y
    systemctl start sshd
    chpasswd <<< 'root:123456'
}

do_test_deploy() {
    log_info "running cypress in e2e"
    cd "${A2_ROOT_DIR}/e2e" || return 1
    export CYPRESS_SKIP_SSO=true
    export CYPRESS_BASE_URL="https://$CONTAINER_HOSTNAME"

    npm install # get dependencies defined in e2e/package.json
    if ! npm run cypress:run:proxy; then
        buildkite-agent artifact upload "cypress/videos/*;cypress/videos/**/*;cypress/screenshots/*;cypress/screenshots/**/*"
        return 1
    fi
}
