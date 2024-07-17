## **The Problem**

Some people use `--save` option when installing packages with `npm`. For example, it is common to see the following command:

```
npm install packagename --save

```

What does the `--save` option do? And why do some people never use it?

## **The Solution**

**npm** or **node package manager** is a package manager for node.

When you download a package using the `npm` command, it installs the packages in the `node_modules` folder **and** adds the installed module as a dependency in your `package.json` file.

Prior to [npm 5.0.0](https://blog.npmjs.org/post/161081169345/v500), npm installed the packages in `node_modules` but didn’t add them as a dependency by default.

If you wanted to save the module as a dependency in the `package.json` file, you had to do it using the `--save` or `-S` option.

As of npm 5.0.0, you no longer need to use this option. Now npm saves all installed packages as dependencies by default.

If you want to save a package as a development-only dependency, you can do so by using the `--save-dev` or `-D` flag.

A few other options are also available that allow you to control how a package is saved using `npm install`. You can read about them on the [official docs](https://docs.npmjs.com/cli/v8/commands/npm-install).