Vagrant.configure("2") do |config|
  config.vm.box = "bento/ubuntu-22.04"
  config.vm.hostname = "webshop-vm"

  config.vm.synced_folder ".", "/srv/webshop"
  config.vm.network "forwarded_port", guest: 8080, host: 8080

  config.vm.provider "virtualbox" do |vb|
    vb.memory = 2048
    vb.cpus = 2
  end

  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    DEBIAN_FRONTEND=noninteractive apt-get install -y ansible
  SHELL

  config.vm.provision "ansible_local" do |ansible|
    ansible.provisioning_path = "/srv/webshop"
    ansible.playbook = "/srv/webshop/playbook.yml"
    ansible.install = false
  end
end
