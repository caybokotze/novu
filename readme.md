# Novu
> Disclaimer: Will currently only support mysql. Feel free to raise a PR to add anything else.
## Wtf is it?
> _"Novu is a tool that you can use to access your database on a local or staging environment securely without the need to open database ports"_

### Why would you find it useful?

Often I find myself on a production server environment on something like Ubuntu (without a nice UI) trying to debug something on the database and using mysql CLI tool to do it. Novu is just a tool that would allow you to do that with a _"marginally better interface"_ than the cli in whatever OS you use, most likely for linux based environments

### How do you use it?

Install the binary on your OS (instructions will follow) and then navigate to localhost:5000 (Do keep in mind that this port will need to be exposed via apache/nginx or whatever you are using)

Wherever you have installed the novu package you will need to navigate to the config file and make sure that you have setup the connection string for your local mysql environment.

Once you are on localhost:5000, you will (should) see an interface that will allow you to configure some connection strings for the database on that server that you would like to connect to. You can then click on the database you would like to target, and then you will be presented with the UI to write some sql code and execute it and view the result in a nice table format. There it's that simple ;)

### If you would like to support me, please give me more work or donate.

I specialise on mobile / web development. I'm not super cheap, but I really appreciate referrals and am happy to chat to anyone.

You can email me here: caybo@mergedigital.io
You can donate to me here: somelink

### Development steps

- **Create HTTP Server**
    - should respond to server requests on localhost
    - should be able to connect to local running mysql instances
    - should act as a proxy for sql commands
    - should have a built in JWT auth system to handle secure requests from the front end


- **Server Database**
    - there should be a dedicated server database strictly for user registration / storing tokens and validating users.
    

- **Create Client project**
    - the client project should have some UI that caters for the ability to write sql code.
    - the ui should also have the ability to execute sql code and display the results in table format.
    - the ui should display the results in horizontal datatable format