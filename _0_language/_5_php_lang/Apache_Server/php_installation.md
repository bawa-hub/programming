vikramkumar@Vikrams-MacBook-Air Apache_Server % brew install shivammathur/php/php@7.4
Warning: shivammathur/php/php@7.4 has been deprecated because it is a versioned formula!
==> Fetching shivammathur/php/php@7.4
==> Downloading https://ghcr.io/v2/shivammathur/php/php/7.4/manifests/7.4.33_4-1
######################################################################################################################################################## 100.0%
==> Downloading https://ghcr.io/v2/shivammathur/php/php/7.4/blobs/sha256:9c4fd04d8b9629d0662498b8e5d44bbee5a82dd8fa6565948712ecc347232dab
######################################################################################################################################################## 100.0%
==> Installing php@7.4 from shivammathur/php
==> Pouring php@7.4--7.4.33_4.arm64_monterey.bottle.1.tar.gz
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set php_ini /opt/homebrew/etc/php/7.4/php.ini system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set php_dir /opt/homebrew/share/pear@7.4 system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set doc_dir /opt/homebrew/share/pear@7.4/doc system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set ext_dir /opt/homebrew/lib/php/pecl/20190902 system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set bin_dir /opt/homebrew/opt/php@7.4/bin system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set data_dir /opt/homebrew/share/pear@7.4/data system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set cfg_dir /opt/homebrew/share/pear@7.4/cfg system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set www_dir /opt/homebrew/share/pear@7.4/htdocs system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set man_dir /opt/homebrew/share/man system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set test_dir /opt/homebrew/share/pear@7.4/test system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear config-set php_bin /opt/homebrew/opt/php@7.4/bin/php system
==> /opt/homebrew/Cellar/php@7.4/7.4.33_4/bin/pear update-channels
==> Caveats
To enable PHP in Apache add the following to httpd.conf and restart Apache:
LoadModule php7_module /opt/homebrew/opt/php@7.4/lib/httpd/modules/libphp7.so

    <FilesMatch \.php$>
        SetHandler application/x-httpd-php
    </FilesMatch>

Finally, check DirectoryIndex includes index.php
DirectoryIndex index.php index.html

The php.ini and php-fpm.ini file can be found in:
/opt/homebrew/etc/php/7.4/

php@7.4 is keg-only, which means it was not symlinked into /opt/homebrew,
because this is an alternate version of another formula.

If you need to have php@7.4 first in your PATH, run:
echo 'export PATH="/opt/homebrew/opt/php@7.4/bin:$PATH"' >> ~/.zshrc
  echo 'export PATH="/opt/homebrew/opt/php@7.4/sbin:$PATH"' >> ~/.zshrc

For compilers to find php@7.4 you may need to set:
export LDFLAGS="-L/opt/homebrew/opt/php@7.4/lib"
export CPPFLAGS="-I/opt/homebrew/opt/php@7.4/include"

To start shivammathur/php/php@7.4 now and restart at login:
brew services start shivammathur/php/php@7.4
Or, if you don't want/need a background service you can just run:
/opt/homebrew/opt/php@7.4/sbin/php-fpm --nodaemonize
==> Summary
🍺 /opt/homebrew/Cellar/php@7.4/7.4.33_4: 499 files, 72.9MB
==> Running `brew cleanup php@7.4`...
Disable this behaviour by setting HOMEBREW_NO_INSTALL_CLEANUP.
Hide these hints with HOMEBREW_NO_ENV_HINTS (see `man brew`).
