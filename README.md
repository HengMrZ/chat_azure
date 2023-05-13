# chat_azure
有用户管理功能的Azure Open AI  API透明代理后台

Azure Open AI API Transparent Proxy Backend

This project is a transparent proxy server for the Azure Open AI API, with user management functions added. It is based on two existing Github projects:

• libli/chat
• haibbo/cf-openai-azure-proxy

This combination of technologies allows users to access the Azure Open AI API securely through the transparent proxy, while also enabling user management functions like authentication and authorization.

Features

• Transparent proxy server for the Azure Open AI API
• User management functions for authentication and authorization

Installation

To install and use this project, you will need to have the following:

• Azure account
• Node.js and NPM installed
• MongoDB installed

Once you have these prerequisites, you can clone this repository, install dependencies, and start the server:

git clone https://github.com/YOUR_USERNAME/YOUR_REPOSITORY.git
cd YOUR_REPOSITORY
npm install
npm start

Make sure to replace YOUR_USERNAME and YOUR_REPOSITORY with the appropriate values for your Github repository.

Usage

To use the transparent proxy server, simply send requests to the appropriate URL (e.g. http://localhost:3000/openai/classification) and include the appropriate API key and model ID. User management functions can be accessed through the appropriate URLs (e.g. http://localhost:3000/users/register).

Contributing

Please feel free to contribute to this project by submitting pull requests or issues on Github.

License

This project is licensed under the MIT License. See the LICENSE file for more details.
