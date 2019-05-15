# Guide
delete the `template` in filename and modify the config.
```Shell
$ mv service.template.config.json service.config.json
$ vim service.config.json
```

> Recommend to use `Envrionment Variable` for password or some secret config. And **Do not upload those sensitive data into github**.