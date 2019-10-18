
ip="192.168.0.11"

ping:
	ping ${ip}

ifconfig:
	vagrant ssh -- -t 'ifconfig'

manual-ssh:
	ssh -i ./.vagrant/machines/default/virtualbox/private_key  vagrant@${ip}
