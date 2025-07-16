<?php

// Authentication Interface
interface AuthenticationProvider
{
    public function authenticate($credentials);
}

// Email/Password Authentication class implementing the interface
class EmailPasswordAuthentication implements AuthenticationProvider
{
    public function authenticate($credentials)
    {
        // Authentication logic for email/password
        echo "Authenticated using email/password.\n";
    }
}

// Social Media Authentication class implementing the interface
class SocialMediaAuthentication implements AuthenticationProvider
{
    private $socialMedia;

    public function __construct($socialMedia)
    {
        $this->socialMedia = $socialMedia;
    }

    public function authenticate($credentials)
    {
        // Authentication logic for social media
        echo "Authenticated using $this->socialMedia credentials.\n";
    }
}

// Single Sign-On (SSO) Authentication class implementing the interface
class SSOAuthentication implements AuthenticationProvider
{
    public function authenticate($credentials)
    {
        // Authentication logic for Single Sign-On
        echo "Authenticated using Single Sign-On.\n";
    }
}

// User Authentication class using polymorphism
class UserAuthentication
{
    private $authenticationProvider;

    public function setAuthenticationProvider(AuthenticationProvider $provider)
    {
        $this->authenticationProvider = $provider;
    }

    public function login($credentials)
    {
        $this->authenticationProvider->authenticate($credentials);
    }
}

// Client code using polymorphism to switch between authentication providers
$userAuth = new UserAuthentication();

// Email/Password authentication
$emailPasswordAuth = new EmailPasswordAuthentication();
$userAuth->setAuthenticationProvider($emailPasswordAuth);
$userAuth->login(["email" => "user@example.com", "password" => "secret"]);

// Social Media authentication
$socialMediaAuth = new SocialMediaAuthentication("Facebook");
$userAuth->setAuthenticationProvider($socialMediaAuth);
$userAuth->login(["token" => "fb_token"]);

// Single Sign-On authentication
$ssoAuth = new SSOAuthentication();
$userAuth->setAuthenticationProvider($ssoAuth);
$userAuth->login(["sso_token" => "sso_token"]);


// The AuthenticationProvider interface defines a common method authenticate that all authentication providers must implement.
// Concrete classes (EmailPasswordAuthentication, SocialMediaAuthentication, SSOAuthentication) implement this interface with their specific authentication logic.
// The UserAuthentication class uses polymorphism by allowing the dynamic switching of authentication providers at runtime through the setAuthenticationProvider method