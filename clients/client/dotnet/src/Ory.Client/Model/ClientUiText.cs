/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v0.0.1-alpha.5
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */


using System;
using System.Collections;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.IO;
using System.Runtime.Serialization;
using System.Text;
using System.Text.RegularExpressions;
using Newtonsoft.Json;
using Newtonsoft.Json.Converters;
using Newtonsoft.Json.Linq;
using System.ComponentModel.DataAnnotations;
using OpenAPIDateConverter = Ory.Client.Client.OpenAPIDateConverter;

namespace Ory.Client.Model
{
    /// <summary>
    /// ClientUiText
    /// </summary>
    [DataContract(Name = "uiText")]
    public partial class ClientUiText : IEquatable<ClientUiText>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUiText" /> class.
        /// </summary>
        [JsonConstructorAttribute]
        protected ClientUiText()
        {
            this.AdditionalProperties = new Dictionary<string, object>();
        }
        /// <summary>
        /// Initializes a new instance of the <see cref="ClientUiText" /> class.
        /// </summary>
        /// <param name="context">The message&#39;s context. Useful when customizing messages..</param>
        /// <param name="id">id (required).</param>
        /// <param name="text">The message text. Written in american english. (required).</param>
        /// <param name="type">type (required).</param>
        public ClientUiText(Object context = default(Object), long id = default(long), string text = default(string), string type = default(string))
        {
            this.Id = id;
            // to ensure "text" is required (not null)
            this.Text = text ?? throw new ArgumentNullException("text is a required property for ClientUiText and cannot be null");
            // to ensure "type" is required (not null)
            this.Type = type ?? throw new ArgumentNullException("type is a required property for ClientUiText and cannot be null");
            this.Context = context;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The message&#39;s context. Useful when customizing messages.
        /// </summary>
        /// <value>The message&#39;s context. Useful when customizing messages.</value>
        [DataMember(Name = "context", EmitDefaultValue = false)]
        public Object Context { get; set; }

        /// <summary>
        /// Gets or Sets Id
        /// </summary>
        [DataMember(Name = "id", IsRequired = true, EmitDefaultValue = false)]
        public long Id { get; set; }

        /// <summary>
        /// The message text. Written in american english.
        /// </summary>
        /// <value>The message text. Written in american english.</value>
        [DataMember(Name = "text", IsRequired = true, EmitDefaultValue = false)]
        public string Text { get; set; }

        /// <summary>
        /// Gets or Sets Type
        /// </summary>
        [DataMember(Name = "type", IsRequired = true, EmitDefaultValue = false)]
        public string Type { get; set; }

        /// <summary>
        /// Gets or Sets additional properties
        /// </summary>
        [JsonExtensionData]
        public IDictionary<string, object> AdditionalProperties { get; set; }

        /// <summary>
        /// Returns the string presentation of the object
        /// </summary>
        /// <returns>String presentation of the object</returns>
        public override string ToString()
        {
            var sb = new StringBuilder();
            sb.Append("class ClientUiText {\n");
            sb.Append("  Context: ").Append(Context).Append("\n");
            sb.Append("  Id: ").Append(Id).Append("\n");
            sb.Append("  Text: ").Append(Text).Append("\n");
            sb.Append("  Type: ").Append(Type).Append("\n");
            sb.Append("  AdditionalProperties: ").Append(AdditionalProperties).Append("\n");
            sb.Append("}\n");
            return sb.ToString();
        }

        /// <summary>
        /// Returns the JSON string presentation of the object
        /// </summary>
        /// <returns>JSON string presentation of the object</returns>
        public virtual string ToJson()
        {
            return Newtonsoft.Json.JsonConvert.SerializeObject(this, Newtonsoft.Json.Formatting.Indented);
        }

        /// <summary>
        /// Returns true if objects are equal
        /// </summary>
        /// <param name="input">Object to be compared</param>
        /// <returns>Boolean</returns>
        public override bool Equals(object input)
        {
            return this.Equals(input as ClientUiText);
        }

        /// <summary>
        /// Returns true if ClientUiText instances are equal
        /// </summary>
        /// <param name="input">Instance of ClientUiText to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(ClientUiText input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Context == input.Context ||
                    (this.Context != null &&
                    this.Context.Equals(input.Context))
                ) && 
                (
                    this.Id == input.Id ||
                    this.Id.Equals(input.Id)
                ) && 
                (
                    this.Text == input.Text ||
                    (this.Text != null &&
                    this.Text.Equals(input.Text))
                ) && 
                (
                    this.Type == input.Type ||
                    (this.Type != null &&
                    this.Type.Equals(input.Type))
                )
                && (this.AdditionalProperties.Count == input.AdditionalProperties.Count && !this.AdditionalProperties.Except(input.AdditionalProperties).Any());
        }

        /// <summary>
        /// Gets the hash code
        /// </summary>
        /// <returns>Hash code</returns>
        public override int GetHashCode()
        {
            unchecked // Overflow is fine, just wrap
            {
                int hashCode = 41;
                if (this.Context != null)
                    hashCode = hashCode * 59 + this.Context.GetHashCode();
                hashCode = hashCode * 59 + this.Id.GetHashCode();
                if (this.Text != null)
                    hashCode = hashCode * 59 + this.Text.GetHashCode();
                if (this.Type != null)
                    hashCode = hashCode * 59 + this.Type.GetHashCode();
                if (this.AdditionalProperties != null)
                    hashCode = hashCode * 59 + this.AdditionalProperties.GetHashCode();
                return hashCode;
            }
        }

        /// <summary>
        /// To validate all properties of the instance
        /// </summary>
        /// <param name="validationContext">Validation context</param>
        /// <returns>Validation Result</returns>
        IEnumerable<System.ComponentModel.DataAnnotations.ValidationResult> IValidatableObject.Validate(ValidationContext validationContext)
        {
            yield break;
        }
    }

}