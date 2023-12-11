<?php
/**
 * UpdateLoginFlowBody
 *
 * PHP version 7.3
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 */

/**
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers.
 *
 * The version of the OpenAPI document: v1.4.6
 * Contact: support@ory.sh
 * Generated by: https://openapi-generator.tech
 * OpenAPI Generator version: 5.4.0
 */

/**
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

namespace Ory\Client\Model;

use \ArrayAccess;
use \Ory\Client\ObjectSerializer;

/**
 * UpdateLoginFlowBody Class Doc Comment
 *
 * @category Class
 * @package  Ory\Client
 * @author   OpenAPI Generator team
 * @link     https://openapi-generator.tech
 * @implements \ArrayAccess<TKey, TValue>
 * @template TKey int|null
 * @template TValue mixed|null
 */
class UpdateLoginFlowBody implements ModelInterface, ArrayAccess, \JsonSerializable
{
    public const DISCRIMINATOR = 'method';

    /**
      * The original name of the model.
      *
      * @var string
      */
    protected static $openAPIModelName = 'updateLoginFlowBody';

    /**
      * Array of property to type mappings. Used for (de)serialization
      *
      * @var string[]
      */
    protected static $openAPITypes = [
        'csrfToken' => 'string',
        'identifier' => 'string',
        'method' => 'string',
        'password' => 'string',
        'passwordIdentifier' => 'string',
        'idToken' => 'string',
        'idTokenNonce' => 'string',
        'provider' => 'string',
        'traits' => 'object',
        'upstreamParameters' => 'object',
        'totpCode' => 'string',
        'webauthnLogin' => 'string',
        'lookupSecret' => 'string',
        'code' => 'string',
        'resend' => 'string'
    ];

    /**
      * Array of property to format mappings. Used for (de)serialization
      *
      * @var string[]
      * @phpstan-var array<string, string|null>
      * @psalm-var array<string, string|null>
      */
    protected static $openAPIFormats = [
        'csrfToken' => null,
        'identifier' => null,
        'method' => null,
        'password' => null,
        'passwordIdentifier' => null,
        'idToken' => null,
        'idTokenNonce' => null,
        'provider' => null,
        'traits' => null,
        'upstreamParameters' => null,
        'totpCode' => null,
        'webauthnLogin' => null,
        'lookupSecret' => null,
        'code' => null,
        'resend' => null
    ];

    /**
     * Array of property to type mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPITypes()
    {
        return self::$openAPITypes;
    }

    /**
     * Array of property to format mappings. Used for (de)serialization
     *
     * @return array
     */
    public static function openAPIFormats()
    {
        return self::$openAPIFormats;
    }

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @var string[]
     */
    protected static $attributeMap = [
        'csrfToken' => 'csrf_token',
        'identifier' => 'identifier',
        'method' => 'method',
        'password' => 'password',
        'passwordIdentifier' => 'password_identifier',
        'idToken' => 'id_token',
        'idTokenNonce' => 'id_token_nonce',
        'provider' => 'provider',
        'traits' => 'traits',
        'upstreamParameters' => 'upstream_parameters',
        'totpCode' => 'totp_code',
        'webauthnLogin' => 'webauthn_login',
        'lookupSecret' => 'lookup_secret',
        'code' => 'code',
        'resend' => 'resend'
    ];

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @var string[]
     */
    protected static $setters = [
        'csrfToken' => 'setCsrfToken',
        'identifier' => 'setIdentifier',
        'method' => 'setMethod',
        'password' => 'setPassword',
        'passwordIdentifier' => 'setPasswordIdentifier',
        'idToken' => 'setIdToken',
        'idTokenNonce' => 'setIdTokenNonce',
        'provider' => 'setProvider',
        'traits' => 'setTraits',
        'upstreamParameters' => 'setUpstreamParameters',
        'totpCode' => 'setTotpCode',
        'webauthnLogin' => 'setWebauthnLogin',
        'lookupSecret' => 'setLookupSecret',
        'code' => 'setCode',
        'resend' => 'setResend'
    ];

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @var string[]
     */
    protected static $getters = [
        'csrfToken' => 'getCsrfToken',
        'identifier' => 'getIdentifier',
        'method' => 'getMethod',
        'password' => 'getPassword',
        'passwordIdentifier' => 'getPasswordIdentifier',
        'idToken' => 'getIdToken',
        'idTokenNonce' => 'getIdTokenNonce',
        'provider' => 'getProvider',
        'traits' => 'getTraits',
        'upstreamParameters' => 'getUpstreamParameters',
        'totpCode' => 'getTotpCode',
        'webauthnLogin' => 'getWebauthnLogin',
        'lookupSecret' => 'getLookupSecret',
        'code' => 'getCode',
        'resend' => 'getResend'
    ];

    /**
     * Array of attributes where the key is the local name,
     * and the value is the original name
     *
     * @return array
     */
    public static function attributeMap()
    {
        return self::$attributeMap;
    }

    /**
     * Array of attributes to setter functions (for deserialization of responses)
     *
     * @return array
     */
    public static function setters()
    {
        return self::$setters;
    }

    /**
     * Array of attributes to getter functions (for serialization of requests)
     *
     * @return array
     */
    public static function getters()
    {
        return self::$getters;
    }

    /**
     * The original name of the model.
     *
     * @return string
     */
    public function getModelName()
    {
        return self::$openAPIModelName;
    }


    /**
     * Associative array for storing property values
     *
     * @var mixed[]
     */
    protected $container = [];

    /**
     * Constructor
     *
     * @param mixed[] $data Associated array of property values
     *                      initializing the model
     */
    public function __construct(array $data = null)
    {
        $this->container['csrfToken'] = $data['csrfToken'] ?? null;
        $this->container['identifier'] = $data['identifier'] ?? null;
        $this->container['method'] = $data['method'] ?? null;
        $this->container['password'] = $data['password'] ?? null;
        $this->container['passwordIdentifier'] = $data['passwordIdentifier'] ?? null;
        $this->container['idToken'] = $data['idToken'] ?? null;
        $this->container['idTokenNonce'] = $data['idTokenNonce'] ?? null;
        $this->container['provider'] = $data['provider'] ?? null;
        $this->container['traits'] = $data['traits'] ?? null;
        $this->container['upstreamParameters'] = $data['upstreamParameters'] ?? null;
        $this->container['totpCode'] = $data['totpCode'] ?? null;
        $this->container['webauthnLogin'] = $data['webauthnLogin'] ?? null;
        $this->container['lookupSecret'] = $data['lookupSecret'] ?? null;
        $this->container['code'] = $data['code'] ?? null;
        $this->container['resend'] = $data['resend'] ?? null;

        // Initialize discriminator property with the model name.
        $this->container['method'] = static::$openAPIModelName;
    }

    /**
     * Show all the invalid properties with reasons.
     *
     * @return array invalid properties with reasons
     */
    public function listInvalidProperties()
    {
        $invalidProperties = [];

        if ($this->container['csrfToken'] === null) {
            $invalidProperties[] = "'csrfToken' can't be null";
        }
        if ($this->container['identifier'] === null) {
            $invalidProperties[] = "'identifier' can't be null";
        }
        if ($this->container['method'] === null) {
            $invalidProperties[] = "'method' can't be null";
        }
        if ($this->container['password'] === null) {
            $invalidProperties[] = "'password' can't be null";
        }
        if ($this->container['provider'] === null) {
            $invalidProperties[] = "'provider' can't be null";
        }
        if ($this->container['totpCode'] === null) {
            $invalidProperties[] = "'totpCode' can't be null";
        }
        if ($this->container['lookupSecret'] === null) {
            $invalidProperties[] = "'lookupSecret' can't be null";
        }
        return $invalidProperties;
    }

    /**
     * Validate all the properties in the model
     * return true if all passed
     *
     * @return bool True if all properties are valid
     */
    public function valid()
    {
        return count($this->listInvalidProperties()) === 0;
    }


    /**
     * Gets csrfToken
     *
     * @return string
     */
    public function getCsrfToken()
    {
        return $this->container['csrfToken'];
    }

    /**
     * Sets csrfToken
     *
     * @param string $csrfToken CSRFToken is the anti-CSRF token
     *
     * @return self
     */
    public function setCsrfToken($csrfToken)
    {
        $this->container['csrfToken'] = $csrfToken;

        return $this;
    }

    /**
     * Gets identifier
     *
     * @return string
     */
    public function getIdentifier()
    {
        return $this->container['identifier'];
    }

    /**
     * Sets identifier
     *
     * @param string $identifier Identifier is the code identifier The identifier requires that the user has already completed the registration or settings with code flow.
     *
     * @return self
     */
    public function setIdentifier($identifier)
    {
        $this->container['identifier'] = $identifier;

        return $this;
    }

    /**
     * Gets method
     *
     * @return string
     */
    public function getMethod()
    {
        return $this->container['method'];
    }

    /**
     * Sets method
     *
     * @param string $method Method should be set to \"code\" when logging in using the code strategy.
     *
     * @return self
     */
    public function setMethod($method)
    {
        $this->container['method'] = $method;

        return $this;
    }

    /**
     * Gets password
     *
     * @return string
     */
    public function getPassword()
    {
        return $this->container['password'];
    }

    /**
     * Sets password
     *
     * @param string $password The user's password.
     *
     * @return self
     */
    public function setPassword($password)
    {
        $this->container['password'] = $password;

        return $this;
    }

    /**
     * Gets passwordIdentifier
     *
     * @return string|null
     */
    public function getPasswordIdentifier()
    {
        return $this->container['passwordIdentifier'];
    }

    /**
     * Sets passwordIdentifier
     *
     * @param string|null $passwordIdentifier Identifier is the email or username of the user trying to log in. This field is deprecated!
     *
     * @return self
     */
    public function setPasswordIdentifier($passwordIdentifier)
    {
        $this->container['passwordIdentifier'] = $passwordIdentifier;

        return $this;
    }

    /**
     * Gets idToken
     *
     * @return string|null
     */
    public function getIdToken()
    {
        return $this->container['idToken'];
    }

    /**
     * Sets idToken
     *
     * @param string|null $idToken IDToken is an optional id token provided by an OIDC provider  If submitted, it is verified using the OIDC provider's public key set and the claims are used to populate the OIDC credentials of the identity. If the OIDC provider does not store additional claims (such as name, etc.) in the IDToken itself, you can use the `traits` field to populate the identity's traits. Note, that Apple only includes the users email in the IDToken.  Supported providers are Apple
     *
     * @return self
     */
    public function setIdToken($idToken)
    {
        $this->container['idToken'] = $idToken;

        return $this;
    }

    /**
     * Gets idTokenNonce
     *
     * @return string|null
     */
    public function getIdTokenNonce()
    {
        return $this->container['idTokenNonce'];
    }

    /**
     * Sets idTokenNonce
     *
     * @param string|null $idTokenNonce IDTokenNonce is the nonce, used when generating the IDToken. If the provider supports nonce validation, the nonce will be validated against this value and required.
     *
     * @return self
     */
    public function setIdTokenNonce($idTokenNonce)
    {
        $this->container['idTokenNonce'] = $idTokenNonce;

        return $this;
    }

    /**
     * Gets provider
     *
     * @return string
     */
    public function getProvider()
    {
        return $this->container['provider'];
    }

    /**
     * Sets provider
     *
     * @param string $provider The provider to register with
     *
     * @return self
     */
    public function setProvider($provider)
    {
        $this->container['provider'] = $provider;

        return $this;
    }

    /**
     * Gets traits
     *
     * @return object|null
     */
    public function getTraits()
    {
        return $this->container['traits'];
    }

    /**
     * Sets traits
     *
     * @param object|null $traits The identity traits. This is a placeholder for the registration flow.
     *
     * @return self
     */
    public function setTraits($traits)
    {
        $this->container['traits'] = $traits;

        return $this;
    }

    /**
     * Gets upstreamParameters
     *
     * @return object|null
     */
    public function getUpstreamParameters()
    {
        return $this->container['upstreamParameters'];
    }

    /**
     * Sets upstreamParameters
     *
     * @param object|null $upstreamParameters UpstreamParameters are the parameters that are passed to the upstream identity provider.  These parameters are optional and depend on what the upstream identity provider supports. Supported parameters are: `login_hint` (string): The `login_hint` parameter suppresses the account chooser and either pre-fills the email box on the sign-in form, or selects the proper session. `hd` (string): The `hd` parameter limits the login/registration process to a Google Organization, e.g. `mycollege.edu`. `prompt` (string): The `prompt` specifies whether the Authorization Server prompts the End-User for reauthentication and consent, e.g. `select_account`.
     *
     * @return self
     */
    public function setUpstreamParameters($upstreamParameters)
    {
        $this->container['upstreamParameters'] = $upstreamParameters;

        return $this;
    }

    /**
     * Gets totpCode
     *
     * @return string
     */
    public function getTotpCode()
    {
        return $this->container['totpCode'];
    }

    /**
     * Sets totpCode
     *
     * @param string $totpCode The TOTP code.
     *
     * @return self
     */
    public function setTotpCode($totpCode)
    {
        $this->container['totpCode'] = $totpCode;

        return $this;
    }

    /**
     * Gets webauthnLogin
     *
     * @return string|null
     */
    public function getWebauthnLogin()
    {
        return $this->container['webauthnLogin'];
    }

    /**
     * Sets webauthnLogin
     *
     * @param string|null $webauthnLogin Login a WebAuthn Security Key  This must contain the ID of the WebAuthN connection.
     *
     * @return self
     */
    public function setWebauthnLogin($webauthnLogin)
    {
        $this->container['webauthnLogin'] = $webauthnLogin;

        return $this;
    }

    /**
     * Gets lookupSecret
     *
     * @return string
     */
    public function getLookupSecret()
    {
        return $this->container['lookupSecret'];
    }

    /**
     * Sets lookupSecret
     *
     * @param string $lookupSecret The lookup secret.
     *
     * @return self
     */
    public function setLookupSecret($lookupSecret)
    {
        $this->container['lookupSecret'] = $lookupSecret;

        return $this;
    }

    /**
     * Gets code
     *
     * @return string|null
     */
    public function getCode()
    {
        return $this->container['code'];
    }

    /**
     * Sets code
     *
     * @param string|null $code Code is the 6 digits code sent to the user
     *
     * @return self
     */
    public function setCode($code)
    {
        $this->container['code'] = $code;

        return $this;
    }

    /**
     * Gets resend
     *
     * @return string|null
     */
    public function getResend()
    {
        return $this->container['resend'];
    }

    /**
     * Sets resend
     *
     * @param string|null $resend Resend is set when the user wants to resend the code
     *
     * @return self
     */
    public function setResend($resend)
    {
        $this->container['resend'] = $resend;

        return $this;
    }
    /**
     * Returns true if offset exists. False otherwise.
     *
     * @param integer $offset Offset
     *
     * @return boolean
     */
    public function offsetExists($offset)
    {
        return isset($this->container[$offset]);
    }

    /**
     * Gets offset.
     *
     * @param integer $offset Offset
     *
     * @return mixed|null
     */
    public function offsetGet($offset)
    {
        return $this->container[$offset] ?? null;
    }

    /**
     * Sets value based on offset.
     *
     * @param int|null $offset Offset
     * @param mixed    $value  Value to be set
     *
     * @return void
     */
    public function offsetSet($offset, $value)
    {
        if (is_null($offset)) {
            $this->container[] = $value;
        } else {
            $this->container[$offset] = $value;
        }
    }

    /**
     * Unsets offset.
     *
     * @param integer $offset Offset
     *
     * @return void
     */
    public function offsetUnset($offset)
    {
        unset($this->container[$offset]);
    }

    /**
     * Serializes the object to a value that can be serialized natively by json_encode().
     * @link https://www.php.net/manual/en/jsonserializable.jsonserialize.php
     *
     * @return mixed Returns data which can be serialized by json_encode(), which is a value
     * of any type other than a resource.
     */
    public function jsonSerialize()
    {
       return ObjectSerializer::sanitizeForSerialization($this);
    }

    /**
     * Gets the string presentation of the object
     *
     * @return string
     */
    public function __toString()
    {
        return json_encode(
            ObjectSerializer::sanitizeForSerialization($this),
            JSON_PRETTY_PRINT
        );
    }

    /**
     * Gets a header-safe presentation of the object
     *
     * @return string
     */
    public function toHeaderValue()
    {
        return json_encode(ObjectSerializer::sanitizeForSerialization($this));
    }
}


