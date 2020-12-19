# go-ddd

Golang Domain Driven Design code structure skeleton.


## Thoughts

It would be great if we can separate golang domains (user, comment) into individual folders, but there are some disadvantages:
- circular dependency. Say if we have a `/user` and `/comment` folder, and we want to add a method to get user comments for both folder. This would not work due to circular dependency. The solution is to put all the interfaces/entity into a single folder called `/domain`
- naming. It is hard to come up with a good naming for packages. And using domain names as package name is not really recommended due to naming collision (`user` as a package name is not good, so is `userimpl`, `usersvc`, `users`).
- common mistake is that we often add the interface in the same folder as the implementation. It is preferable to place all the interfaces in a separate folder, and the implementation in another folder. The interface folder act as a contract. Whenever we want to split the service/refactor/rewrite, just remove the implementation folder.


## Should controller be separated?

It is probably better to put the controller in a separate folder, outside of the implementation. This is because the controller is dependent on the framework used, and we want to keep that layer separate from the implementation folder. Also, controller belongs to the `transporter` category. There can be other ways of transport such as `grpc`, `graphql` etc that just calls the `service` layer.

## cmd folder

The applications (cli or server) are placed in the `cmd` folder. Each app is a composition of different use cases, and it depends on the usage. We can run an api server to allow clients to call the api, as well as running scripts that allows admins to interact with the application directly. Note that it is not wrong to place the main login to run the server in the root. It is just a matter of preference.
