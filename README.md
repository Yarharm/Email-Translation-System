# Automatic Email Translation

Automatic email translation is a project inspired by the challenges that Minister of Imigration & Citizenship is facing. It aims to help all new immigrants with the initial language barrier.

The main idea is to setup an automatic translation of the incoming email correspondance to any preferred language.

Hopefuly, the following tool can be supplied to all enrolled in [Cours de français à temps complet](https://www.immigration-quebec.gouv.qc.ca/fr/langue-francaise/temps-complet/index.html) as a part of the tool package.

### Prerequisites

* Gmail account which serves as a source of information for the tool.

* Avialable server to host Jenkins for automation purposes.

### Installing

1. Obtain [Yandex Tanslation API](https://tech.yandex.com/translate/) credentials to get access to translation service. Do not rename or change location of the obtained `translation_token.json`.

2. Setup configuration file: `config.json`

|  Property | Description  |
|---|---|
| translationRecipient  | Gmail account of the translated message recipient  |   
| subject  | Desired subject for the emails recieved from the tool  |  
|  translationLanguage | Preferred language for translation   |   

```javascript
{
  translationRecipient: "recipient_email@gmail.com",
  subject: "Translation Bot",
  translationLanguage: "en-ru"
}
```
Refer to the list of supported languages at [Supported languages](https://tech.yandex.com/translate/doc/dg/concepts/api-overview-docpage/).

3. Execute `make` command.

```golang
make
```
The first time you would need to grant access to your gmail account.
Do not rename or change location of the obtained `token.json` and `credentials.json`.

**Gmail authentication scope is set to the lowest sensitivity, targeting only Primary Inbox unread emails.**

*Authentication will no longer be necessary after the initial setup is finished.*

4. At this point the tool can be utilized locally for development or on-demand translation. For the automatization, proceed with the following steps:

    1. Install [Docker](https://hub.docker.com/search?type=edition&offering=community) on your machine.

    2. Refer to the following [Jenkins setup with Docker](https://jenkins.io/doc/book/installing/) for your desired OS.

    3. Add [Poll Mailbox Trigger Plugin](https://wiki.jenkins.io/display/JENKINS/Poll+Mailbox+Trigger+Plugin) that would trigger out tool on every incoming inbox message.

    4. Add [Go plugin](https://plugins.jenkins.io/golang/) for the code execution.

    5. Add [Authentication Tokens API Plugin](https://plugins.jenkins.io/authentication-tokens/) to supply all the tokens received during authorization process.

    6. Follow the [Freestyle job guide](https://www.guru99.com/create-builds-jenkins-freestyle-project.html) and point your Jenkins job towards your repository. At this point everything should be up and running, since Jenkinsfile is already present.

