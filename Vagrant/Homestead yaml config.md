
```yaml
authorize: ~/.ssh/id_rsa_homestead.pub

keys:
    - ~/.ssh/github
    - ~/.ssh/gitlab

folders:
    - map: C:/Users/china/working/
      to: /home/vagrant/working

sites:
    - map: homestead.test
      to: /home/vagrant/code/public
```