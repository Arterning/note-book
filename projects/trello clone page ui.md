
## dash board

![[Pasted image 20231219150959.png]]



## create board
![[Pasted image 20231219151142.png]]

here we can choose tie board image and create board 

detail code in `FormPopover`

the image picker use grid layout, and detail code in `FormPicker`

```js
{
	images.map((image) => {
		<div key={image.id}
		     onClick={()=> {
			     setSelecteImage(image.id)
		     }}>		
			<Image/>
			{
				selectedImageId === image.id 
				&& <ImgeSelected/>
			}
		</div>
	})
}
```

use state to store the image user selected

`const [selectedImageId, setSelectedImageId] = useState(null);`


we just need to create board by titile and image, the rest thing we can built it later
```js
  const onSubmit = (formData: FormData) => {

    const title = formData.get("title") as string;

    const image = formData.get("image") as string;

    execute({ title, image });

  }
```


## server action
we use server action to mutate server data

```js
const { execute, fieldErrors } = useAction(createBoard, {
    onSuccess: (data) => {

      toast.success("Board created!");

      closeRef.current?.click();

      router.push(`/board/${data.id}`);

    },

    onError: (error) => {

      toast.error(error);

      proModal.onOpen();

    }

  });
```




```js
import { useState, useCallback } from "react";

  

import { ActionState, FieldErrors } from "@/lib/create-safe-action";

  

type Action<TInput, TOutput> = (data: TInput) => Promise<ActionState<TInput, TOutput>>;

  

interface UseActionOptions<TOutput> {

  onSuccess?: (data: TOutput) => void;

  onError?: (error: string) => void;

  onComplete?: () => void;

};

  

/**

 * Custom useAction hook

 * "use client" component can use this hook on the client side

 * @param action

 * @param options

 */

export const useAction = <TInput, TOutput> (

  action: Action<TInput, TOutput>,

  options: UseActionOptions<TOutput> = {}

) => {

  const [fieldErrors, setFieldErrors] = useState<FieldErrors<TInput> | undefined>(

    undefined

  );

  const [error, setError] = useState<string | undefined>(undefined);

  const [data, setData] = useState<TOutput | undefined>(undefined);

  const [isLoading, setIsLoading] = useState<boolean>(false);

  

  const execute = useCallback(

    async (input: TInput) => {

      setIsLoading(true);

  

      try {

        const result = await action(input);

  

        if (!result) {

          return;

        }

  

        setFieldErrors(result.fieldErrors);

  

        if (result.error) {

          setError(result.error);

          options.onError?.(result.error);

        }

  

        if (result.data) {

          setData(result.data);

          options.onSuccess?.(result.data);

        }

      } finally {

        setIsLoading(false);

        options.onComplete?.();

      }

    },

    [action, options]

  );

  

  return {

    execute,

    fieldErrors,

    error,

    data,

    isLoading,

  };

};
```




