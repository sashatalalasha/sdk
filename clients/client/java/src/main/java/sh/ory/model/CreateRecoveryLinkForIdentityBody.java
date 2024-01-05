/*
 * Ory APIs
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.4.9
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
 * Create Recovery Link for Identity Request Body
 */
@ApiModel(description = "Create Recovery Link for Identity Request Body")
@javax.annotation.Generated(value = "org.openapitools.codegen.languages.JavaClientCodegen", date = "2024-01-05T11:29:05.592825961Z[Etc/UTC]")
public class CreateRecoveryLinkForIdentityBody {
  public static final String SERIALIZED_NAME_EXPIRES_IN = "expires_in";
  @SerializedName(SERIALIZED_NAME_EXPIRES_IN)
  private String expiresIn;

  public static final String SERIALIZED_NAME_IDENTITY_ID = "identity_id";
  @SerializedName(SERIALIZED_NAME_IDENTITY_ID)
  private String identityId;

  public CreateRecoveryLinkForIdentityBody() {
  }

  public CreateRecoveryLinkForIdentityBody expiresIn(String expiresIn) {
    
    this.expiresIn = expiresIn;
    return this;
  }

   /**
   * Link Expires In  The recovery link will expire after that amount of time has passed. Defaults to the configuration value of &#x60;selfservice.methods.code.config.lifespan&#x60;.
   * @return expiresIn
  **/
  @javax.annotation.Nullable
  @ApiModelProperty(value = "Link Expires In  The recovery link will expire after that amount of time has passed. Defaults to the configuration value of `selfservice.methods.code.config.lifespan`.")

  public String getExpiresIn() {
    return expiresIn;
  }


  public void setExpiresIn(String expiresIn) {
    this.expiresIn = expiresIn;
  }


  public CreateRecoveryLinkForIdentityBody identityId(String identityId) {
    
    this.identityId = identityId;
    return this;
  }

   /**
   * Identity to Recover  The identity&#39;s ID you wish to recover.
   * @return identityId
  **/
  @javax.annotation.Nonnull
  @ApiModelProperty(required = true, value = "Identity to Recover  The identity's ID you wish to recover.")

  public String getIdentityId() {
    return identityId;
  }


  public void setIdentityId(String identityId) {
    this.identityId = identityId;
  }

  /**
   * A container for additional, undeclared properties.
   * This is a holder for any undeclared properties as specified with
   * the 'additionalProperties' keyword in the OAS document.
   */
  private Map<String, Object> additionalProperties;

  /**
   * Set the additional (undeclared) property with the specified name and value.
   * If the property does not already exist, create it otherwise replace it.
   *
   * @param key name of the property
   * @param value value of the property
   * @return the CreateRecoveryLinkForIdentityBody instance itself
   */
  public CreateRecoveryLinkForIdentityBody putAdditionalProperty(String key, Object value) {
    if (this.additionalProperties == null) {
        this.additionalProperties = new HashMap<String, Object>();
    }
    this.additionalProperties.put(key, value);
    return this;
  }

  /**
   * Return the additional (undeclared) property.
   *
   * @return a map of objects
   */
  public Map<String, Object> getAdditionalProperties() {
    return additionalProperties;
  }

  /**
   * Return the additional (undeclared) property with the specified name.
   *
   * @param key name of the property
   * @return an object
   */
  public Object getAdditionalProperty(String key) {
    if (this.additionalProperties == null) {
        return null;
    }
    return this.additionalProperties.get(key);
  }


  @Override
  public boolean equals(Object o) {
    if (this == o) {
      return true;
    }
    if (o == null || getClass() != o.getClass()) {
      return false;
    }
    CreateRecoveryLinkForIdentityBody createRecoveryLinkForIdentityBody = (CreateRecoveryLinkForIdentityBody) o;
    return Objects.equals(this.expiresIn, createRecoveryLinkForIdentityBody.expiresIn) &&
        Objects.equals(this.identityId, createRecoveryLinkForIdentityBody.identityId)&&
        Objects.equals(this.additionalProperties, createRecoveryLinkForIdentityBody.additionalProperties);
  }

  @Override
  public int hashCode() {
    return Objects.hash(expiresIn, identityId, additionalProperties);
  }

  @Override
  public String toString() {
    StringBuilder sb = new StringBuilder();
    sb.append("class CreateRecoveryLinkForIdentityBody {\n");
    sb.append("    expiresIn: ").append(toIndentedString(expiresIn)).append("\n");
    sb.append("    identityId: ").append(toIndentedString(identityId)).append("\n");
    sb.append("    additionalProperties: ").append(toIndentedString(additionalProperties)).append("\n");
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
    openapiFields.add("expires_in");
    openapiFields.add("identity_id");

    // a set of required properties/fields (JSON key names)
    openapiRequiredFields = new HashSet<String>();
    openapiRequiredFields.add("identity_id");
  }

 /**
  * Validates the JSON Object and throws an exception if issues found
  *
  * @param jsonObj JSON Object
  * @throws IOException if the JSON Object is invalid with respect to CreateRecoveryLinkForIdentityBody
  */
  public static void validateJsonObject(JsonObject jsonObj) throws IOException {
      if (jsonObj == null) {
        if (!CreateRecoveryLinkForIdentityBody.openapiRequiredFields.isEmpty()) { // has required fields but JSON object is null
          throw new IllegalArgumentException(String.format("The required field(s) %s in CreateRecoveryLinkForIdentityBody is not found in the empty JSON string", CreateRecoveryLinkForIdentityBody.openapiRequiredFields.toString()));
        }
      }

      // check to make sure all required properties/fields are present in the JSON string
      for (String requiredField : CreateRecoveryLinkForIdentityBody.openapiRequiredFields) {
        if (jsonObj.get(requiredField) == null) {
          throw new IllegalArgumentException(String.format("The required field `%s` is not found in the JSON string: %s", requiredField, jsonObj.toString()));
        }
      }
      if ((jsonObj.get("expires_in") != null && !jsonObj.get("expires_in").isJsonNull()) && !jsonObj.get("expires_in").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `expires_in` to be a primitive type in the JSON string but got `%s`", jsonObj.get("expires_in").toString()));
      }
      if (!jsonObj.get("identity_id").isJsonPrimitive()) {
        throw new IllegalArgumentException(String.format("Expected the field `identity_id` to be a primitive type in the JSON string but got `%s`", jsonObj.get("identity_id").toString()));
      }
  }

  public static class CustomTypeAdapterFactory implements TypeAdapterFactory {
    @SuppressWarnings("unchecked")
    @Override
    public <T> TypeAdapter<T> create(Gson gson, TypeToken<T> type) {
       if (!CreateRecoveryLinkForIdentityBody.class.isAssignableFrom(type.getRawType())) {
         return null; // this class only serializes 'CreateRecoveryLinkForIdentityBody' and its subtypes
       }
       final TypeAdapter<JsonElement> elementAdapter = gson.getAdapter(JsonElement.class);
       final TypeAdapter<CreateRecoveryLinkForIdentityBody> thisAdapter
                        = gson.getDelegateAdapter(this, TypeToken.get(CreateRecoveryLinkForIdentityBody.class));

       return (TypeAdapter<T>) new TypeAdapter<CreateRecoveryLinkForIdentityBody>() {
           @Override
           public void write(JsonWriter out, CreateRecoveryLinkForIdentityBody value) throws IOException {
             JsonObject obj = thisAdapter.toJsonTree(value).getAsJsonObject();
             obj.remove("additionalProperties");
             // serialize additonal properties
             if (value.getAdditionalProperties() != null) {
               for (Map.Entry<String, Object> entry : value.getAdditionalProperties().entrySet()) {
                 if (entry.getValue() instanceof String)
                   obj.addProperty(entry.getKey(), (String) entry.getValue());
                 else if (entry.getValue() instanceof Number)
                   obj.addProperty(entry.getKey(), (Number) entry.getValue());
                 else if (entry.getValue() instanceof Boolean)
                   obj.addProperty(entry.getKey(), (Boolean) entry.getValue());
                 else if (entry.getValue() instanceof Character)
                   obj.addProperty(entry.getKey(), (Character) entry.getValue());
                 else {
                   obj.add(entry.getKey(), gson.toJsonTree(entry.getValue()).getAsJsonObject());
                 }
               }
             }
             elementAdapter.write(out, obj);
           }

           @Override
           public CreateRecoveryLinkForIdentityBody read(JsonReader in) throws IOException {
             JsonObject jsonObj = elementAdapter.read(in).getAsJsonObject();
             validateJsonObject(jsonObj);
             // store additional fields in the deserialized instance
             CreateRecoveryLinkForIdentityBody instance = thisAdapter.fromJsonTree(jsonObj);
             for (Map.Entry<String, JsonElement> entry : jsonObj.entrySet()) {
               if (!openapiFields.contains(entry.getKey())) {
                 if (entry.getValue().isJsonPrimitive()) { // primitive type
                   if (entry.getValue().getAsJsonPrimitive().isString())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsString());
                   else if (entry.getValue().getAsJsonPrimitive().isNumber())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsNumber());
                   else if (entry.getValue().getAsJsonPrimitive().isBoolean())
                     instance.putAdditionalProperty(entry.getKey(), entry.getValue().getAsBoolean());
                   else
                     throw new IllegalArgumentException(String.format("The field `%s` has unknown primitive type. Value: %s", entry.getKey(), entry.getValue().toString()));
                 } else if (entry.getValue().isJsonArray()) {
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), List.class));
                 } else { // JSON object
                     instance.putAdditionalProperty(entry.getKey(), gson.fromJson(entry.getValue(), HashMap.class));
                 }
               }
             }
             return instance;
           }

       }.nullSafe();
    }
  }

 /**
  * Create an instance of CreateRecoveryLinkForIdentityBody given an JSON string
  *
  * @param jsonString JSON string
  * @return An instance of CreateRecoveryLinkForIdentityBody
  * @throws IOException if the JSON string is invalid with respect to CreateRecoveryLinkForIdentityBody
  */
  public static CreateRecoveryLinkForIdentityBody fromJson(String jsonString) throws IOException {
    return JSON.getGson().fromJson(jsonString, CreateRecoveryLinkForIdentityBody.class);
  }

 /**
  * Convert an instance of CreateRecoveryLinkForIdentityBody to an JSON string
  *
  * @return JSON string
  */
  public String toJson() {
    return JSON.getGson().toJson(this);
  }
}

