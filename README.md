# go-ddd

Golang Domain Driven Design code structure skeleton.


## Thoughts

It would be great if we can separate golang domains (user, comment) into individual folders, but there are some disadvantages:
- circular dependency. Say if we have a `/user` and `/comment` folder, and we want to add a method to get user comments for both folder. This would not work due to circular dependency. The solution is to put all the interfaces/entity into a single folder called `/domain`
- naming. It is hard to come up with a good naming for packages. And using domain names as package name is not really recommended due to naming collision (`user` as a package name is not good, so is `userimpl`, `usersvc`, `users`).
- common mistake is that we often add the interface in the same folder as the implementation. It is preferable to place all the interfaces in a separate folder, and the implementation in another folder. The interface folder act as a contract. Whenever we want to split the service/refactor/rewrite, just remove the implementation folder.
