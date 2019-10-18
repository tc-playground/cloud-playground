exec {"apt-get update":
  command => "/usr/bin/apt-get update",
}

package {"apache2":
  require => Exec["apt-get update"],
}

file { "/var/www": 
  ensure => link, 
  target => "/vagrant", 
  force => true,
}

exec {"make index.html":
  command => 'echo "<strong>Apache Vagrant Machine -" `hostname` "</strong>" > /var/www/index.html',
  path => "/bin/"
}
