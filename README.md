# jaip-auth-lambda

This repo is provides a lambda that handles redirection from the JSTOR login pages to the JAIP platform's admin page. Its resonsibility is two-fold. 
1. Handle the JSTOR uuid cookie so that it can be picked up by the JAIP frontend as a URL fragment
1. Redirect the browser to whichever login page the user came from (this could be `admin.pep.jstor.org`, `admin.test-pep.jstor.org`, or an admin page in the ephemeral environment).

## Deployment

Deployment is handled via Serverless v3.

### Requirements

1. [Install Serverless v3](https://www.serverless.com/framework/docs) using `sudo npm i serverless@3 -g`
1. Create a `secrets.yml` file in the root directory of the project. The required values are defined in `secrets.example.yml`. 
    * In order to determine the appropriate values for the subnets and security group, the most straightforward method would be to check the existing Lambdas, `labs-pep-auth-test-proxy` and `labs-pep-auth-prod-proxy`.(these currently share the same subnets and security group, though they are in different environments).
    * For the `environment` value, use the environment to which you plan to deploy (either `test` or `prod`).
    * For the role, use either `test-standard-lambda-role` or `prod-standard-lambda-role`, again depending on which environment you are deploying to.

### Deployment

Deployment is done via the Makefile in the rood directory of the project. Once the requirements are satisfied, proceed with the following steps:

1. If any local builds exist, run `make clean`.
1. Build and deploy with `stage-test`.
1. Verify that the Lambda works by visiting `admin.test-pep.jstor.org` and completing the login process.
    * Click the log in button on the landing page.
    * Enter a username and password on the JSTOR login page (note that the email should have at least reviewer privileges for at least one group at test-pep.jstor.org).
    * Verify that you've been redirected to `admin.test-pep.jstor.org` and are now logged in as an admin user. If there is still a Log in button in the navigation menu, you have not successfully logged in. You should see your name at the top of the navigation menu.
1. Deploy to prod with `stage-prod`.