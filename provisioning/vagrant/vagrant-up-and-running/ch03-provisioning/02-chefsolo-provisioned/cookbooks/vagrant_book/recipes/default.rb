execute "apt-get update" 
package "apache2"
execute "rm -rf /var/www" 
link "/var/www" do
	to "/vagrant" 
end
execute 'echo "<strong>Apache Vagrant Machine -" `hostname` "</strong>" > /var/www/index.html'
