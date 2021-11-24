## Introduction

One of the many ways of creating incidents at Rootly is through our Slack bot. Our customers rely on this bot to perform various actions as declaring, changing incident status, paging on-call people without leaving Slack. 

## Project

This project goal is to develop a simple Slack bot and web UI.

- Slack command can receive 2 commands
    - /rootly declare <title>
    - /rootly resolve
- Web UI
    - List incidents created

### /rootly declare <title>

- This command will display a dialog to create a new incident
    - This command works in any channel.
    - A title is required, a description is optional, a severity (sev0,sev1,sev2) is optional.
    - New incident will create a dedicated Slack channel so people can troubleshoot the issue.
    - You will store the new incident created in an `incidents` table backed by PostgreSQL.
    - You will create a simple UI to display list of incidents with their title, description, severity and the creator.

### /rootly resolve

- This command works in a dedicated incident slack channel only.
- Only available in the dedicated Slack incident channel, this will mark the status of the incident as `resolved` and display back to the channel the time it took the team to get to resolution.

### Deployment ( Version 1 )

There are obviously numerous ways to deploy this service and expose it on the internet. This is project is very small and straightforward, i had like you to deploy it on [Heroku](https://heroku.com/) or [Render](https://render.com) using their free plan.

### Deployment ( Version 2, only for SRE role )

There are obviously numerous ways to deploy this service and expose it on the internet. Although this is project is very small and straightforward, I’d like for you to deploy it scalably.

For this project, Dockerize the application and use whatever technology (e.g. Kubernetes, AWS ECS, Terraform, etc.) you think is best.

### Notes

These are open-ended instructions—you can take it in whatever direction you want. How you display information and UI, in general, is also up to you. In the interest of time, please do not spend too much time on the UI.

In the spirit of transparency and out of respect for your time, I want you know why doing this project is helpful for the both of us:

- This is a more realistic, and, hopefully, more fun way for me to get a feel for how you work and think.
- You will get experience with our whole stack. You will be working with Rails and PostgreSQL a lot with Rootly in the short-term.
- You will be going through the Slack documentation to complete the task.
- You will be able to use Docker with Rails and deploy scalably in production. This may be one of your earliest projects with us, so hopefully you can use the learnings from this project and apply it to our code. ( Only for SRE role )
- I’ll be able to get a sense for your coding style and organization.
- I will be able to get you involved in imperative Rootly projects as soon as you start and get you to hit the ground running.

### Test

- You will need to allow distribution so i can install your slack app in a test workspace of mine. ( You don't need to publish your app and wait for Slack approval )
    
    ![https://s3-us-west-2.amazonaws.com/secure.notion-static.com/40a85723-44b1-4db8-84f8-93e03bf09e4d/Screen_Shot_2021-05-06_at_10.58.43_AM.png](https://s3-us-west-2.amazonaws.com/secure.notion-static.com/40a85723-44b1-4db8-84f8-93e03bf09e4d/Screen_Shot_2021-05-06_at_10.58.43_AM.png)
    
- Add [https://github.com/kwent](https://github.com/kwent) (`kwentakill@gmail.com`) as collaborator on GitHub so i can see your code.

## Here is some stuff that may be helpful.

Slack

- [https://api.slack.com/start](https://api.slack.com/start)
- [https://api.slack.com/interactivity/slash-commands](https://api.slack.com/interactivity/slash-commands)
- [https://ngrok.com/download](https://ngrok.com/download) ( To help testing your callbacks locally )

Heroku

- [https://devcenter.heroku.com/articles/getting-started-with-rails6](https://devcenter.heroku.com/articles/getting-started-with-rails6)

Render

- [https://render.com/docs/deploy-rails](https://render.com/docs/deploy-rails)