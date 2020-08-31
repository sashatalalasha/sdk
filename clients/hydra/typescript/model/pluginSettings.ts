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
import { PluginDevice } from './pluginDevice';
import { PluginMount } from './pluginMount';

export class PluginSettings {
    /**
    * args
    */
    'args': Array<string>;
    /**
    * devices
    */
    'devices': Array<PluginDevice>;
    /**
    * env
    */
    'env': Array<string>;
    /**
    * mounts
    */
    'mounts': Array<PluginMount>;

    static discriminator: string | undefined = undefined;

    static attributeTypeMap: Array<{name: string, baseName: string, type: string}> = [
        {
            "name": "args",
            "baseName": "Args",
            "type": "Array<string>"
        },
        {
            "name": "devices",
            "baseName": "Devices",
            "type": "Array<PluginDevice>"
        },
        {
            "name": "env",
            "baseName": "Env",
            "type": "Array<string>"
        },
        {
            "name": "mounts",
            "baseName": "Mounts",
            "type": "Array<PluginMount>"
        }    ];

    static getAttributeTypeMap() {
        return PluginSettings.attributeTypeMap;
    }
}

