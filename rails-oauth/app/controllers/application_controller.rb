class ApplicationController < ActionController::Base
  before_action :authenticate_user!  # Ensure the user is authenticated before any action
end
