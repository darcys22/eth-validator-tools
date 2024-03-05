#!/bin/bash

for file in ./validator_keys/*
do
  lighthouse --network goerli account validator exit --no-confirmation --keystore "$file" --password-file ./pass
done
