/*
 * Ory APIs
 *
 * Documentation for all public and administrative Ory APIs. Administrative APIs can only be accessed with a valid Personal Access Token. Public APIs are mostly used in browsers. 
 *
 * The version of the OpenAPI document: v1.1.24
 * Contact: support@ory.sh
 * Generated by: https://github.com/openapitools/openapi-generator.git
 */

using System;
using System.IO;
using System.Collections.Generic;
using System.Collections.ObjectModel;
using System.Linq;
using System.Reflection;
using RestSharp;
using Xunit;

using Ory.Client.Client;
using Ory.Client.Api;
// uncomment below to import models
//using Ory.Client.Model;

namespace Ory.Client.Test.Api
{
    /// <summary>
    ///  Class for testing ProjectApi
    /// </summary>
    /// <remarks>
    /// This file is automatically generated by OpenAPI Generator (https://openapi-generator.tech).
    /// Please update the test case below to test the API endpoint.
    /// </remarks>
    public class ProjectApiTests : IDisposable
    {
        private ProjectApi instance;

        public ProjectApiTests()
        {
            instance = new ProjectApi();
        }

        public void Dispose()
        {
            // Cleanup when everything is done.
        }

        /// <summary>
        /// Test an instance of ProjectApi
        /// </summary>
        [Fact]
        public void InstanceTest()
        {
            // TODO uncomment below to test 'IsType' ProjectApi
            //Assert.IsType<ProjectApi>(instance);
        }

        /// <summary>
        /// Test CreateProject
        /// </summary>
        [Fact]
        public void CreateProjectTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //ClientCreateProjectBody clientCreateProjectBody = null;
            //var response = instance.CreateProject(clientCreateProjectBody);
            //Assert.IsType<ClientProject>(response);
        }

        /// <summary>
        /// Test CreateProjectApiKey
        /// </summary>
        [Fact]
        public void CreateProjectApiKeyTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string project = null;
            //ClientInlineObject3 clientInlineObject3 = null;
            //var response = instance.CreateProjectApiKey(project, clientInlineObject3);
            //Assert.IsType<ClientProjectApiKey>(response);
        }

        /// <summary>
        /// Test DeleteProjectApiKey
        /// </summary>
        [Fact]
        public void DeleteProjectApiKeyTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string project = null;
            //string tokenId = null;
            //instance.DeleteProjectApiKey(project, tokenId);
        }

        /// <summary>
        /// Test GetActiveProjectInConsole
        /// </summary>
        [Fact]
        public void GetActiveProjectInConsoleTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //var response = instance.GetActiveProjectInConsole();
            //Assert.IsType<ClientActiveProjectInConsole>(response);
        }

        /// <summary>
        /// Test GetProject
        /// </summary>
        [Fact]
        public void GetProjectTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string projectId = null;
            //var response = instance.GetProject(projectId);
            //Assert.IsType<ClientProject>(response);
        }

        /// <summary>
        /// Test GetProjectMembers
        /// </summary>
        [Fact]
        public void GetProjectMembersTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string projectId = null;
            //var response = instance.GetProjectMembers(projectId);
            //Assert.IsType<List<ClientCloudAccount>>(response);
        }

        /// <summary>
        /// Test ListProjectApiKeys
        /// </summary>
        [Fact]
        public void ListProjectApiKeysTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string project = null;
            //var response = instance.ListProjectApiKeys(project);
            //Assert.IsType<List<ClientProjectApiKey>>(response);
        }

        /// <summary>
        /// Test ListProjects
        /// </summary>
        [Fact]
        public void ListProjectsTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //var response = instance.ListProjects();
            //Assert.IsType<List<ClientProjectMetadata>>(response);
        }

        /// <summary>
        /// Test PatchProject
        /// </summary>
        [Fact]
        public void PatchProjectTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string projectId = null;
            //List<ClientJsonPatch> clientJsonPatch = null;
            //var response = instance.PatchProject(projectId, clientJsonPatch);
            //Assert.IsType<ClientSuccessfulProjectUpdate>(response);
        }

        /// <summary>
        /// Test PurgeProject
        /// </summary>
        [Fact]
        public void PurgeProjectTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string projectId = null;
            //instance.PurgeProject(projectId);
        }

        /// <summary>
        /// Test RemoveProjectMember
        /// </summary>
        [Fact]
        public void RemoveProjectMemberTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string projectId = null;
            //string memberId = null;
            //instance.RemoveProjectMember(projectId, memberId);
        }

        /// <summary>
        /// Test SetActiveProjectInConsole
        /// </summary>
        [Fact]
        public void SetActiveProjectInConsoleTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //ClientSetActiveProjectInConsoleBody clientSetActiveProjectInConsoleBody = null;
            //instance.SetActiveProjectInConsole(clientSetActiveProjectInConsoleBody);
        }

        /// <summary>
        /// Test SetProject
        /// </summary>
        [Fact]
        public void SetProjectTest()
        {
            // TODO uncomment below to test the method and replace null with proper value
            //string projectId = null;
            //ClientSetProject clientSetProject = null;
            //var response = instance.SetProject(projectId, clientSetProject);
            //Assert.IsType<ClientSuccessfulProjectUpdate>(response);
        }
    }
}
