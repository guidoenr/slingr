#!/bin/bash

appDir=$(realpath "$(dirname "$0")")
enginesDir=$(cd ../.. && pwd)


# enginesFile=$enginesDir/engines/aiware/transcribe/speechmatics-chunk/cicd/cicd-engines.json
enginesFile=$appDir/engines.json
tmpFile=$appDir/temp.json


getEnginesToUpdate() {
  declare -A EnginesMap

  # Get all engines from JSON file and store it in temporary file
  $(jq -r '.engines' "$enginesFile" > $tmpFile)

  # Get keys with are the engine IDs
  keys=$(jq -r '. | keys | .[]' $tmpFile)

  # Loop through each engine and select only the ones which we want to update the version to newer one
  for key in $keys;do
    value=$(jq -r ". | values[\"$key\"] | .smVersion == \"sm901\"" $tmpFile)
    if $value == true
    then

      # Get the name of the engine and replace it with newer version
      name=$(jq -r ". | values[\"$key\"] | .engineName" $tmpFile)

      toReplace="(v9.0.1)"
      replaceWith="(v9.1.0)"

      newEngineName=${name/$toReplace/"$replaceWith"}

      # Store key(engineID) and value(new version engine name)
      EnginesMap[$key]=$newEngineName
      
      # echo $key
      # echo $newEngineName

      url="https://api.stage.us-1.veritone.com/v3/graphql"
      TOKEN="f4f05d32-0ea6-4484-b840-9c2a3304bc0f"
      

      # @guido: pasas los dos argumentos asi a la func asi los puede leer, el segundo como tiene espacios
      # va con "$varName" y el primero no
      updateEngineName $key "$newEngineName"

      # curl -g \
      # -X POST \
      # -H "Content-Type: application/json" \
      # -H "Authorization: Bearer $TOKEN" \
      # --data '{ "query": "mutation { updateEngine(input: {id:\'$key'\, name: \"'"$newEngineName"'\"}){id\n name\n}}" }' \
      # $url

    fi
  done
  
  # echo "LENGTH: ${#EnginesMap[*]}"
  # for key in "${!EnginesMap[@]}"; do
  #   echo "$key ${EnginesMap[$key]}"
  # done
  # Remove temp file once done
  rm $tmpFile
} 

# le tendria q pasar engineID y tambien el idioma (o hacer query para obtener info del nombre actual)
updateEngineName(){
  url="https://api.stage.us-1.veritone.com/v3/graphql"
  TOKEN="0c835f78-78f6-4ff9-ade2-bd688a2f4cda"
  key="1e5b4e36-acbe-49f1-bebf-1d3a0edd6219"
  name="Google - OCR"

  # aca accedes a los args que le pasas segun el orden en que son pasados
  # $1 = key, $2 = newEngineName

  curl -g \
  -X POST \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $TOKEN" \
  --data '{ "query": "mutation { updateEngine(input: {id: \"'$1'\", name: \"'"$2"'\"}){id\n name\n}}" }' \
  $url
}


# Do GraphQL update
# Update Engine.json file with new smVersion

# FUNCIONES:  

getEnginesToUpdate
# updateEngineName