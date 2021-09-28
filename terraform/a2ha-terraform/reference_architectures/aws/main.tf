module "aws" {
  ami_filter_name                    = var.ami_filter_name
  ami_filter_owner                   = var.ami_filter_owner
  ami_filter_virt_type               = var.ami_filter_virt_type
  automate_ebs_volume_iops           = var.automate_ebs_volume_iops
  automate_ebs_volume_size           = var.automate_ebs_volume_size
  automate_ebs_volume_type           = var.automate_ebs_volume_type
  automate_fqdn                      = var.automate_fqdn
  automate_instance_count            = var.automate_instance_count
  automate_lb_certificate_arn        = var.automate_lb_certificate_arn
  automate_server_instance_type      = var.automate_server_instance_type
  aws_ami_id                         = var.aws_ami_id
  aws_profile                        = var.aws_profile
  aws_region                         = var.aws_region
  aws_ssh_key_file                   = var.ssh_key_file
  aws_ssh_key_pair_name              = var.aws_ssh_key_pair_name
  aws_ssh_user                       = var.ssh_user
  chef_ebs_volume_iops               = var.automate_ebs_volume_iops
  chef_ebs_volume_size               = var.automate_ebs_volume_size
  chef_ebs_volume_type               = var.automate_ebs_volume_type
  chef_server_instance_count         = var.chef_server_instance_count
  chef_server_instance_type          = var.chef_server_instance_type
  chef_server_lb_certificate_arn     = var.chef_server_lb_certificate_arn
  elasticsearch_ebs_volume_iops      = var.elasticsearch_ebs_volume_iops
  elasticsearch_ebs_volume_size      = var.elasticsearch_ebs_volume_size
  elasticsearch_ebs_volume_type      = var.elasticsearch_ebs_volume_type
  elasticsearch_instance_count       = var.elasticsearch_instance_count
  elasticsearch_listen_port          = var.elasticsearch_listen_port
  elasticsearch_server_instance_type = var.elasticsearch_server_instance_type
  pgleaderchk_listen_port            = var.pgleaderchk_listen_port
  postgresql_ebs_volume_iops         = var.postgresql_ebs_volume_iops
  postgresql_ebs_volume_size         = var.postgresql_ebs_volume_size
  postgresql_ebs_volume_type         = var.postgresql_ebs_volume_type
  postgresql_instance_count          = var.postgresql_instance_count
  postgresql_listen_port             = var.postgresql_listen_port
  postgresql_server_instance_type    = var.postgresql_server_instance_type
  proxy_listen_port                  = var.proxy_listen_port
  source                             = "./modules/aws"
  tags                               = var.aws_tags
}

module "aws-output" {
  source                    = "./modules/aws_output"
  automate_public_ips       = module.aws.automate_public_ips
  automate_private_ips      = module.aws.automate_private_ips
  chef_server_public_ips    = module.aws.chef_server_public_ips
  chef_server_private_ips   = module.aws.chef_server_private_ips
  postgresql_public_ips     = module.aws.postgresql_public_ips
  postgresql_private_ips    = module.aws.postgresql_private_ips
  elasticsearch_public_ips  = module.aws.elasticsearch_public_ips
  elasticsearch_private_ips = module.aws.elasticsearch_private_ips
  automate_fqdn             = module.aws.automate_fqdn
  automate_frontend_urls    = module.aws.automate_frontend_urls
}

module "aws-destroy" {
  source = "./modules/aws_destroy"
}