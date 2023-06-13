# Install Node

Created time: June 7, 2023 1:34 PM
Tags: node

# How to Install Node.js and NVM

Node.js is a JavaScript runtime built on Chrome's V8 JavaScript engine. It allows developers to run JavaScript on the server-side. NVM (Node Version Manager) is a tool that allows you to manage multiple versions of Node.js.

## Installing NVM

1. Open your terminal or command prompt.
2. Install NVM by running the following command:
    
    ```
    curl -O https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh | bash
    
    ```
    
    Alternatively, you can install NVM using the wget command:
    
    ```
    wget -qO- <https://raw.githubusercontent.com/nvm-sh/nvm/v0.39.0/install.sh> | bash
    
    ```
    
3. Verify that NVM has been installed correctly by running:
    
    ```
    nvm --version
    
    ```
    

## Installing Node.js

1. Open your terminal or command prompt.
2. Install the latest version of Node.js by running the following command:
    
    ```
    nvm install node
    
    ```
    
    If you want to install a specific version of Node.js, you can run:
    
    ```
    nvm install {version number}
    
    ```
    
    For example, to install Node.js version 14.17.0, you would run:
    
    ```
    nvm install 14.17.0
    
    ```
    
3. Verify that Node.js has been installed correctly by running:
    
    ```
    node --version
    
    ```
    

Congratulations! You have successfully installed Node.js and NVM on your system.