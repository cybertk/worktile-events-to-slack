# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure(2) do |config|
  config.vm.box = "coverit/golang-dev"
  config.vm.box_version = "20150512.0.0"

  config.vm.provision "shell", privileged: false, inline: <<-SHELL
    # Setup golang workspace
    mkdir -p ~/go/src/github.com/cybertk
    ln -s /vagrant ~/go/src/github.com/cybertk/worktile-events-to-slack

    # Audo reload with gin
    GOPATH=~/go go get github.com/codegangsta/gin
  SHELL

    if File.exists?("script/custom-vagrant")
        config.vm.provision "shell", path: "script/custom-vagrant"
    end
end
