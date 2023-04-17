# -*- mode: ruby -*-
# vi: set ft=ruby :

# All Vagrant configuration is done below. The "2" in Vagrant.configure
# configures the configuration version (we support older styles for
# backwards compatibility). Please don't change it unless you know what
# you're doing.

$as_root = <<AS_ROOT
echo Running root script...
sudo apt-get -y update
sudo apt-get -y install vim curl git wget build-essential
wget https://dl.google.com/go/go1.15.linux-amd64.tar.gz
sha256sum go1.15.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.15.linux-amd64.tar.gz
rm -f go1.15.linux-amd64.tar.gz
chown -R vagrant:vagrant go
echo Root script complete.
AS_ROOT

$as_vagrant = <<AS_VAGRANT
echo Running vagrant script...
echo 'export GOPATH="$HOME/go"' >> ~/.bashrc
echo 'export GONOSUMDB="github.com/deliveryhero/*,github.com/pedidosya/*"' >> ~/.bashrc
export GOPATH="$HOME/go"
echo 'export PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"' >> ~/.bashrc
export PATH="$PATH:/usr/local/go/bin:$GOPATH/bin"
go version
echo Vagrant script complete.
AS_VAGRANT

Vagrant.configure("2") do |config|
    config.vm.box = "hashicorp/bionic64"
    config.vm.network "private_network", ip: "10.0.0.10"
    config.vm.synced_folder ".", "/vagrant", disabled: true
    config.vm.synced_folder "../@project_name@", "/home/vagrant/go/src/github.com/pedidosya/@project_name@"
    config.vm.synced_folder "~/.ssh", "/home/vagrant/.ssh"
    config.vm.provision "file", source: "~/.gitconfig", destination: ".gitconfig"
    config.vm.provision "file", source: "~/.vault_token", destination: ".vault_token"
    config.vm.provision :shell, :inline => $as_root, :privileged => true
    config.vm.provision :shell, :inline => $as_vagrant, :privileged => false
end
