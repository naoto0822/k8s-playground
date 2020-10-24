# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|

  # Master Node１
  config.vm.define "master" do |machine|
    machine.vm.box = "bento/centos-7.7"
    machine.vm.hostname = 'master'
    machine.vm.network :private_network,ip: "192.168.10.10"
    machine.vm.provider "virtualbox" do |vbox|
      vbox.gui = false
      vbox.cpus = 2
      vbox.memory = 1024
    end
    machine.vm.synced_folder "./ansible/sync", "/home/vagrant/sync", owner: "vagrant",
      group: "vagrant", mount_options: ["dmode=777", "fmode=777"]

    machine.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "./ansible/k8s.yml"
      ansible.version = "latest"
      ansible.verbose = false
      ansible.limit = "master"
      ansible.inventory_path = "ansible/hosts"
    end

    machine.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "./ansible/k8s_master.yml"
      ansible.version = "latest"
      ansible.verbose = false
      ansible.install = false
      ansible.limit = "master"
      ansible.inventory_path = "ansible/hosts"
    end
  end

  # Worker Node１
  config.vm.define 'worker1' do |machine|
    machine.vm.box = "bento/centos-7.7"
    machine.vm.hostname = 'worker1'
    machine.vm.network :private_network,ip: "192.168.10.20"
    machine.vm.provider "virtualbox" do |vbox|
      vbox.gui = false
      vbox.cpus = 1
      vbox.memory = 1024
    end
    machine.vm.synced_folder "./ansible/sync", "/home/vagrant/sync", owner: "vagrant",
      group: "vagrant", mount_options: ["dmode=777", "fmode=777"]

    machine.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "./ansible/k8s.yml"
      ansible.version = "latest"
      ansible.verbose = false
      ansible.limit = "worker1"
      ansible.inventory_path = "ansible/hosts"
    end

    machine.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "./ansible/k8s_workers.yml"
      ansible.version = "latest"
      ansible.verbose = false
      ansible.install = false
      ansible.limit = "worker1"
      ansible.inventory_path = "ansible/hosts"
    end
  end

  # Worker Node２
  config.vm.define 'worker2' do |machine|
    machine.vm.box = "bento/centos-7.7"
    machine.vm.hostname = 'worker2'
    machine.vm.network :private_network,ip: "192.168.10.30"
    machine.vm.provider "virtualbox" do |vbox|
      vbox.gui = false
      vbox.cpus = 1
      vbox.memory = 1024
    end
    machine.vm.synced_folder "./ansible/sync", "/home/vagrant/sync", owner: "vagrant",
      group: "vagrant", mount_options: ["dmode=777", "fmode=777"]

    machine.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "./ansible/k8s.yml"
      ansible.version = "latest"
      ansible.verbose = false
      ansible.limit = "worker2"
      ansible.inventory_path = "ansible/hosts"
    end

    machine.vm.provision "ansible_local" do |ansible|
      ansible.playbook = "./ansible/k8s_workers.yml"
      ansible.version = "latest"
      ansible.verbose = false
      ansible.install = false
      ansible.limit = "worker2"
      ansible.inventory_path = "ansible/hosts"
    end
  end
end
