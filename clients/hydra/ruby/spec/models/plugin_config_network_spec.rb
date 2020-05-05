=begin
#ORY Hydra

#Welcome to the ORY Hydra HTTP API documentation. You will find documentation for all HTTP APIs here.

The version of the OpenAPI document: latest

Generated by: https://openapi-generator.tech
OpenAPI Generator version: 4.2.3

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OryHydraClient::PluginConfigNetwork
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'PluginConfigNetwork' do
  before do
    # run before each test
    @instance = OryHydraClient::PluginConfigNetwork.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of PluginConfigNetwork' do
    it 'should create an instance of PluginConfigNetwork' do
      expect(@instance).to be_instance_of(OryHydraClient::PluginConfigNetwork)
    end
  end
  describe 'test attribute "type"' do
    it 'should work' do
      # assertion here. ref: https://www.relishapp.com/rspec/rspec-expectations/docs/built-in-matchers
    end
  end

end