
Search `useQuery`

	in CardModal.jsx, 
```jsx
const { data: cardData } = useQuery<CardWithList>({
    queryKey: ["card", id],
    queryFn: () => fetcher(`/api/cards/${id}`),
  });

  const { data: auditLogsData } = useQuery<AuditLog[]>({
    queryKey: ["card-logs", id],
    queryFn: () => fetcher(`/api/cards/${id}/logs`),
  });
```



create a new card and click , dialog popup

![[Pasted image 20231229100259.png]]


When User Update the Title and Description, the Title Field and Activity Field will updated, because we use react-query here



in the card-modal->header.tsx, we see here to invalidate cache after update completed

```jsx
  const { execute } = useAction(updateCard, {
    onSuccess: (data) => {
      queryClient.invalidateQueries({
        queryKey: ["card", data.id]
      });

      queryClient.invalidateQueries({
        queryKey: ["card-logs", data.id]
      });

      toast.success(`Renamed to "${data.title}"`);
      setTitle(data.title);
    },
    onError: (error) => {
      toast.error(error);
    }
  });
```



