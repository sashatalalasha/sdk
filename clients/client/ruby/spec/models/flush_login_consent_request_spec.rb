=begin
#Ory APIs

#Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 

The version of the OpenAPI document: v0.2.0-alpha.29
Contact: support@ory.sh
Generated by: https://openapi-generator.tech
OpenAPI Generator version: 6.0.1

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryClient::FlushLoginConsentRequest
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OryClient::FlushLoginConsentRequest do
  let(:instance) { OryClient::FlushLoginConsentRequest.new }

  describe 'test an instance of FlushLoginConsentRequest' do
    it 'should create an instance of FlushLoginConsentRequest' do
      expect(instance).to be_instance_of(OryClient::FlushLoginConsentRequest)
    end
  end
  describe 'test attribute "not_after"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end