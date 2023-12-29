## C: Create
If we need to build a form for user to create something, we can use Modal or NewPage, if the form is simple and fields are small, we can use Modal to simply create, However, if the form is big and the fields are big, Consider to navigate user to a new page is a better choise.

After created, we should refresh page and navgate user to list page or detai page

Toast popup was needed

## R: Read
In the past , we often use useEffect to load Data and manage the error, data, isLoading state by ourself, but now we can use react-query and useSWR to fetch data and manage the remote data status

## U: Update

After updated, we should refresh page and navgate user to list page or detai page
There are two design patter
- Provide Update button, user should click update button to update
- Not Provide Update button, user can update data with enter key
Toast popup was needed

## D: Delete

After Deleted, we should refresh page and navgate user to list page or detai page
We should provide a confirm dialog for user to avoid user Delete accidently
Toast popup was needed
