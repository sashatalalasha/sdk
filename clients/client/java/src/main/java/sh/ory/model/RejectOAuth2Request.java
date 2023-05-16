/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.29
 * Contact: support@ory.sh
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


package sh.ory.model;

import java.util.Objects;
import java.util.Arrays;
import com.google.gson.TypeAdapter;
import com.google.gson.annotations.JsonAdapter;
import com.google.gson.annotations.SerializedName;
import com.google.gson.stream.JsonReader;
import com.google.gson.stream.JsonWriter;
import io.swagger.annotations.ApiModel;
import io.swagger.annotations.ApiModelProperty;
import java.io.IOException;

import com.google.gson.Gson;
import com.google.gson.GsonBuilder;
import com.google.gson.JsonArray;
import com.google.gson.JsonDeserializationContext;
import com.google.gson.JsonDeserializer;
import com.google.gson.JsonElement;
import com.google.gson.JsonObject;
import com.google.gson.JsonParseException;
import com.google.gson.TypeAdapterFactory;
import com.google.gson.reflect.TypeToken;

import java.lang.reflect.Type;
import java.util.HashMap;
import java.util.HashSet;
import java.util.List;
import java.util.Map;
import java.util.Map.Entry;
import java.util.Set;

import sh.ory.JSON;

/**
 * RejectOAuth2Request
 */
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2023-05-16T08:08:56.165385030Z[Etc/UTC]")
public class RejectOAuth2Request {
  public static final String SERIALIZED_NAME_ERROR = "error";
  @SerializedName(SERIALIZED_NAME_ERROR)
  private String error;

  public static final String SERIALIZED_NAME_ERROR_DEBUG = "error_debug";
  @SerializedName(SERIALIZED_NAME_ERROR_DEBUG)
  private String errorDebug;

  public static final String SERIALIZED_NAME_ERROR_DESCRIPTION = "error_description";
  @SerializedName(SERIALIZED_NAME_ERROR_DESCRIPTION)
  private String errorDescription;

  public static final String SERIALIZED_NAME_ERROR_HINT = "error_hint";
  @SerializedName(SERIALIZED_NAME_ERROR_HINT)
  private String errorHint;

  public static final String SERIALIZED_NAME_STATUS_CODE = "status_code";
  @SerializedName(SERIALIZED_NAME_STATUS_CODE)
  private Long statusCode;

  public RejectOAuth2Request() {
  }

  public RejectOAuth2Request error(String error) {
    
    this.error = error;
    return this;
  }

   /**
   * The error should follow the OAuth2 error format (e.g. &#x60;invalid_request&#x60;, &#x60;login_required&#x60;).  Defaults to &#x60;request_denied&#x60;.
   * @return error
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "The error should follow the OAuth2 error format (e.g. `invalid_request`, `login_required`).  Defaults to `request_denied`.")

  public String getError() {
    return error;
  }


  public void setError(String error) {
    this.error = error;
  }


  public RejectOAuth2Request errorDebug(String errorDebug) {
    
    this.errorDebug = errorDebug;
    return this;
  }

   /**
   * Debug contains information to help resolve the problem as a developer. Usually not exposed to the public but only in the server logs.
   * @return errorDebug
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Debug contains information to help resolve the problem as a developer. Usually not exposed to the public but only in the server logs.")

  public String getErrorDebug() {
    return errorDebug;
  }


  public void setErrorDebug(String errorDebug) {
    this.errorDebug = errorDebug;
  }


  public RejectOAuth2Request errorDescription(String errorDescription) {
    
    this.errorDescription = errorDescription;
    return this;
  }

   /**
   * Description of the error in a human readable format.
   * @return errorDescription
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Description of the error in a human readable format.")

  public String getErrorDescription() {
    return errorDescription;
  }


  public void setErrorDescription(String errorDescription) {
    this.errorDescription = errorDescription;
  }


  public RejectOAuth2Request errorHint(String errorHint) {
    
    this.errorHint = errorHint;
    return this;
  }

   /**
   * Hint to help resolve the error.
   * @return errorHint
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Hint to help resolve the error.")

  public String getErrorHint() {
    return errorHint;
  }


  public void setErrorHint(String errorHint) {
    this.errorHint = errorHint;
  }


  public RejectOAuth2Request statusCode(Long statusCode) {
    
    this.statusCode = statusCode;
    return this;
  }

   /**
   * Represents the HTTP status code of the error (e.g. 401 or 403)  Defaults to 400
   * @return statusCode
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Represents the HTTP status code of the error (e.g. 401 or 403)  Defaults to 400")

  public Long getStatusCode() {
    return statusCode;
  }


  public void setStatusCode(Long statusCode) {
    this.statusCode = statusCode;
  }



  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    RejectOAuth2Request rejectOAuth2Request = (RejectOAuth2Request) o;
    return Objects.equals(this.error, rejectOAuth2Request.error) &&
        Objects.equals(this.errorDebug, rejectOAuth2Request.errorDebug) &&
        Objects.equals(this.errorDescription, rejectOAuth2Request.errorDescription) &&
        Objects.equals(this.errorHint, rejectOAuth2Request.errorHint) &&
        Objects.equals(this.statusCode, rejectOAuth2Request.statusCode);
  }

  @Override
  public int hashCode() {
    return Objects.hash(error, errorDebug, errorDescription, errorHint, statusCode);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class RejectOAuth2Request {\n");
    sb.append("    error: ").append(toIndentedString(error)).append("\n");
    sb.append("    errorDebug: ").append(toIndentedString(errorDebug)).append("\n");
    sb.append("    errorDescription: ").append(toIndentedString(errorDescription)).append("\n");
    sb.append("    errorHint: ").append(toIndentedString(errorHint)).append("\n");
    sb.append("    statusCode: ").append(toIndentedString(statusCode)).append("\n");
    sb.append("}");
    return sb.toString();
  }

  /**
   * Convert the given object to string with each line indented by 4 spaces
   * (except the first line).
   */
  private String toIndentedString(Object o) {
    if (o == null) {
      return "null";
    }
    return o.toString().replace("\n", "\n    ");
  }


  public static HashSet<String> openapiFields;
  public static HashSet<String> openapiRequiredFields;

  static {
    // a set of all properties/fields (JSON key names)
    openapiFields = new HashSet<String>();
    openapiFields.add("error");
    openapiFields.add("error_debug");
    openapiFields.add("error_description");
    openapiFields.add("error_hint");
    openapiFields.add("status_code");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to RejectOAuth2Request
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!RejectOAuth2Request.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in RejectOAuth2Request is not found in the empty JSON string", RejectOAuth2Request.openapiRequiredFields.toString()));
        }
      }

      Set<Entry<String, JsonElement>> entries = jsonObj.entrySet();
      // check to see if the JSON string contains additional fields
      for (Entry<String, JsonElement> entry : entries) {
        if (!RejectOAuth2Request.openapiFields.contains(entry.getKey())) {
          throw new IllegalArgumentException(String.format("The field `%s` in the JSON string is not defined in the `RejectOAuth2Request` properties. JSON: %s", entry.getKey(), jsonObj.toString()));
        }
      }
      if ((jsonObj.get("error") != null && !jsonObj.get("error").isJsonNull()) && !jsonObj.get("error").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `error` to be a primitive type in the JSON string but got `%s`", jsonObj.get("error").toString()));
      }
      if ((jsonObj.get("error_debug") != null && !jsonObj.get("error_debug").isJsonNull()) && !jsonObj.get("error_debug").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `error_debug` to be a primitive type in the JSON string but got `%s`", jsonObj.get("error_debug").toString()));
      }
      if ((jsonObj.get("error_description") != null && !jsonObj.get("error_description").isJsonNull()) && !jsonObj.get("error_description").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `error_description` to be a primitive type in the JSON string but got `%s`", jsonObj.get("error_description").toString()));
      }
      if ((jsonObj.get("error_hint") != null && !jsonObj.get("error_hint").isJsonNull()) && !jsonObj.get("error_hint").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `error_hint` to be a primitive type in the JSON string but got `%s`", jsonObj.get("error_hint").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!RejectOAuth2Request.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'RejectOAuth2Request' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<RejectOAuth2Request> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(RejectOAuth2Request.class));

       return (TypeAdapter<T>) new TypeAdapter<RejectOAuth2Request>() {
           @Override
           public void write(JsonWriter out, RejectOAuth2Request value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             elementAdapter.write(out, obj);
           }

           @Override
           public RejectOAuth2Request read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             return thisAdapter.fromJsonTree(jsonObj);
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of RejectOAuth2Request given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of RejectOAuth2Request
  * @throws IOException if the JSON string is invalid with respect to RejectOAuth2Request
  */
  public static RejectOAuth2Request fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, RejectOAuth2Request.class);
  }

 /**
  * Convert an instance of RejectOAuth2Request to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

