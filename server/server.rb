require 'sinatra/base'
require 'sinatra/assetpack'

class App < Sinatra::Base
  register Sinatra::AssetPack
  assets do
    serve '/js', :from => 'javascripts'
    serve '/css', :from => 'stylesheets'

    css_compression :sass
  end

  get '/data' do
    hour = Time.now - (86400 * 30)
    output = ""
    while hour < Time.now
      output += "#{hour},#{5.0 * rand()},\n"

      hour += 3600
    end

    output
  end

  get '/' do
    erb :index
  end

  run! if app_file == $0
end
