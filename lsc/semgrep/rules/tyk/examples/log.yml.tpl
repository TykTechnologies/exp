  - id: log.:log:.remove.withField
    pattern: 'log.WithField("prefix",":prefix:")'
    fix: :log:
    languages:
      - go
    message: Replacing :log: prefix with logger (autofix)
    severity: WARNING

  - id: log.:log:.remove.withFields
    pattern: 'log.WithFields(logrus.Fields{"prefix":":prefix:"})'
    fix: :log:
    languages:
      - go
    message: Replacing :log: prefix with logger (autofix)
    severity: WARNING

  - id: log.:log:.remove.prefix.from.Fields
    pattern: 'log.WithFields(logrus.Fields{"prefix":":prefix:",$X})'
    fix: |
      :log:.WithFields(logrus.Fields{
      	$X,
      })
    languages:
      - go
    message: Replacing :log: prefix with logger (autofix)
    severity: WARNING
