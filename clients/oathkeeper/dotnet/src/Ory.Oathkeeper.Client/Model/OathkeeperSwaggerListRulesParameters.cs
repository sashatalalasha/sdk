/*
 * ORY Oathkeeper
 *
 * ORY Oathkeeper is a reverse proxy that checks the HTTP Authorization for validity against a set of rules. This service uses Hydra to validate access tokens and policies.
 *
 * The version of the OpenAPI document: v0.0.0-alpha.54
 * Contact: hi@ory.am
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
using OpenAPIDateConverter = Ory.Oathkeeper.Client.Client.OpenAPIDateConverter;

namespace Ory.Oathkeeper.Client.Model
{
    /// <summary>
    /// SwaggerListRulesParameters swagger list rules parameters
    /// </summary>
    [DataContract(Name = "swaggerListRulesParameters")]
    public partial class OathkeeperSwaggerListRulesParameters : IEquatable<OathkeeperSwaggerListRulesParameters>, IValidatableObject
    {
        /// <summary>
        /// Initializes a new instance of the <see cref="OathkeeperSwaggerListRulesParameters" /> class.
        /// </summary>
        /// <param name="limit">The maximum amount of rules returned. in: query.</param>
        /// <param name="offset">The offset from where to start looking. in: query.</param>
        public OathkeeperSwaggerListRulesParameters(long limit = default(long), long offset = default(long))
        {
            this.Limit = limit;
            this.Offset = offset;
            this.AdditionalProperties = new Dictionary<string, object>();
        }

        /// <summary>
        /// The maximum amount of rules returned. in: query
        /// </summary>
        /// <value>The maximum amount of rules returned. in: query</value>
        [DataMember(Name = "limit", EmitDefaultValue = false)]
        public long Limit { get; set; }

        /// <summary>
        /// The offset from where to start looking. in: query
        /// </summary>
        /// <value>The offset from where to start looking. in: query</value>
        [DataMember(Name = "offset", EmitDefaultValue = false)]
        public long Offset { get; set; }

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
            sb.Append("class OathkeeperSwaggerListRulesParameters {\n");
            sb.Append("  Limit: ").Append(Limit).Append("\n");
            sb.Append("  Offset: ").Append(Offset).Append("\n");
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
            return this.Equals(input as OathkeeperSwaggerListRulesParameters);
        }

        /// <summary>
        /// Returns true if OathkeeperSwaggerListRulesParameters instances are equal
        /// </summary>
        /// <param name="input">Instance of OathkeeperSwaggerListRulesParameters to be compared</param>
        /// <returns>Boolean</returns>
        public bool Equals(OathkeeperSwaggerListRulesParameters input)
        {
            if (input == null)
                return false;

            return 
                (
                    this.Limit == input.Limit ||
                    this.Limit.Equals(input.Limit)
                ) && 
                (
                    this.Offset == input.Offset ||
                    this.Offset.Equals(input.Offset)
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
                hashCode = hashCode * 59 + this.Limit.GetHashCode();
                hashCode = hashCode * 59 + this.Offset.GetHashCode();
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
