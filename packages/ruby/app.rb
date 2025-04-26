require 'sinatra'
require 'yaml'

get '/load' do
  # Unsafe deserialization vulnerability
  data = YAML.load(params[:data])
  data.to_s
end 