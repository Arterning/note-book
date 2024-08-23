
```python
  
  

@router.get('/ocr_rest', summary='OCR迁移',

    dependencies=[

    ],

)

async def ocr_rest():

    docs = await doc_service.get_all_docs()

    k = 0

    for doc in docs:

        await ocr_exist_doc(doc)

        try:

            await handle_ocr_doc(doc)

        except Exception as e:

            traceback.print_exc()

        k += 1

        log.info(f"toal {len(docs)}, now {k}")

    return await response_base.success(data="OK")

  
  
  

async def ocr_exist_doc(doc, file_name):

    loop = asyncio.get_event_loop()

    title = doc.title

    uid = doc.uuid

    file_name = f"{uid}.pdf"

    if doc.ocr_res:

        log.info(f"no need ocr {title}")

    else:

        response = minio_client.get_object(bucket_name, file_name)

        file_content = response.read()

        log.info(f"extract_ocr {title}, before ocr_res is {doc.ocr_res}")

        ocr_res, content = await loop.run_in_executor(None, extrat_content_ocr, file_content)

        await doc_service.base_update(pk=doc.id, obj={

                "ocr_res": ocr_res,

                "content": content,

            })

  
  

async def handle_ocr_doc(doc):

    loop = asyncio.get_event_loop()

    ocr_res = doc.ocr_res

    if not ocr_res:

        return

    title = doc.title

    if doc.location:

        log.info(f"no need location {title}")

    else:

        log.info(f"extract_location {title}")

        location = await loop.run_in_executor(None, extract_location, ocr_res)

    if doc.purpose:

        log.info(f"no need purpose {title}")

    else:

        purpose = await loop.run_in_executor(None, extract_purpose, ocr_res)

        log.info(f"extract_purpose {title}")

  

    if doc.trans:

        log.info(f"no need trans {title}")

    else:

        trans = await loop.run_in_executor(None, translate_document, ocr_res)

        log.info(f"translate_document {title}")

  

    if doc.digest:

        log.info(f"no need digest {title}")

    else:

        digest = await loop.run_in_executor(None, digest_document, ocr_res)

        log.info(f"digest_document {title}")

  

    await doc_service.base_update(pk=doc.id, obj={

                "location" : location,

                "purpose": purpose,

                "trans": trans,

                "digest": digest

            })
```