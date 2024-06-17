Doorkeeper.configure do
  # Change the ORM that doorkeeper will use (needs plugins)
  # Currently supported options are :active_record, :mongoid2, :mongoid3, :mongo_mapper
  orm :active_record

  default_scopes :public
  optional_scopes :user

  # Ensure this block is called when scopes are requested
  grant_flows %w[authorization_code client_credentials password]

  # This block will be called to check whether the resource owner is authenticated or not.
  resource_owner_authenticator do
    current_user || warden.authenticate!(scope: :user)
  end

  # If you want to restrict access to the web interface for adding oauth authorized applications
  admin_authenticator do
    current_user && current_user.admin? || redirect_to(new_user_session_url)
  end

  # Authorization Code expiration time (default 10 minutes).
  authorization_code_expires_in 10.minutes

  # Access token expiration time (default 2 hours).
  access_token_expires_in 2.hours

  # Use a custom class for generating access tokens
  # access_token_generator '::Doorkeeper::JWT'

  # Reuse access tokens for the same resource owner within an application (default false)
  reuse_access_token

  # Issue access tokens with refresh token (default true)
  use_refresh_token

  # Hashing the tokens for added security (default false)
  hash_token_secrets

  # Forces the usage of the HTTPS protocol for all endpoints (default true)
  force_ssl_in_redirect_uri !Rails.env.development?
end
