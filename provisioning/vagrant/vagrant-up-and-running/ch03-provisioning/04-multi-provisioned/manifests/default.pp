exec {"make index.html":
  command => 'echo "<strong>Apache Vagrant Machine -" `hostname` "</strong>" > /var/www/index.html',
  path => "/bin/"
}
