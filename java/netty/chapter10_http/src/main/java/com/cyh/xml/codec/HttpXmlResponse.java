/*
 * Copyright 2013-2018 Lilinfeng.
 *  
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *  
 *      http://www.apache.org/licenses/LICENSE-2.0
 *  
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package com.cyh.xml.codec;

import io.netty.handler.codec.http.FullHttpResponse;
import lombok.Data;

/**
 * @author Administrator
 * @date 2014年3月1日
 * @version 1.0
 */
@Data
public class HttpXmlResponse {

    private FullHttpResponse httpResponse;
    private Object result;

    public HttpXmlResponse(FullHttpResponse httpResponse, Object result) {
        this.httpResponse = httpResponse;
        this.result = result;
    }


}
