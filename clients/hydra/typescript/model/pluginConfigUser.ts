/**
 * ORY Hydra
 * Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.
 *
 * The version of the OpenAPI document: v1.7.3
 * 
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { RequestFile } from '../api';

/**
* PluginConfigUser plugin config user
*/
export class PluginConfigUser {
    /**
    * g ID
    */
    'gID'?: number;
    /**
    * UID
    */
    'uID'?: number;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "gID",
            "baseName": "GID",
            "type": "number"
        },
        {
            "name": "uID",
            "baseName": "UID",
            "type": "number"
        }    ];

    static getAttributeTypeMap() {
        return PluginConfigUser.attributeTypeMap;
    }
}

