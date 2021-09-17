require 'yaml'
require 'ffi'
require 'json'

module Cel
    extend FFI::Library
    ffi_lib File.expand_path("../cel/shared.so", File.dirname(__FILE__))

    attach_function :evaluer, [:string, :string, :string, :string], :string
end

blobs_info = File.read('../dummyblobs.json')

blobs = JSON.parse(blobs_info)

configs = YAML.load_file('../cel/config.yml')
i = 1
blobs.each do |blob|
    puts i

    blob['content'] = File.read(blob['full_path'])

    configs['frameworks'].each do |config|
        puts config['name']
        if blob['frameworks'].include? config['name']
            puts 'config to be applied'
            config['settings'].each do |k|
                puts '--------------------------------------'
                puts blob['name']
                puts k['name']
                puts k['value']
                
                puts 'calling go method'
                        
                puts Cel.evaluer(blob['name'], blob['full_path'], blob['content'], k['value'])
                puts '--------------------------------------'
            end
        end
    end

    i += 1
end