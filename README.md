## my gofiber + templ workbench 

### the associated fetch

When working with forms and relational data, it's often necessary to refetch changed data to inherit client state. When we have a form for a User, where I can change the address by assigning a new AddressID (imagine a combobox search feature where I can select the address from a list of addresses), we have to re-evaluate what the user selected when we render the form.

To reproduce the behavior, follow these steps:
1. Start the server: `go run cmd/server/main.go`
2. Open the browser at `http://localhost:3000/`
3. Click "edit" on Ellen Doe
4. Clear the email field
5. Change the address from "1" to "2"
6. Submit the form
7. The form should now show an error for the email and "Street 2, 4321 Ytic" for the address

I have solved my scenario of this problem using `RenderForm` inside of [handlers.go:37](./pkg/users/handlers.go:37).