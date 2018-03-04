aws lambda create-function --region ap-northeast-1 --function-name GoMovies --zip-file fileb://./deployment.zip --runtime go1.x --role arn:aws:iam::$account_id:role/$role --handler main
