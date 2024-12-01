# async_serverless_application_demo

### api server

- nautilus:
  - 商品一覧の取得や購入処理を行う基盤となるAPI

- ibis:
  - S3向け画像コンテンツのダウンロード/アップロード用の署名付きURLを発行するAPI

- coral:
  - ユーザ情報管理用API (DynamoDBへレコードを投入し、データの永続化を行うAPI)

### batch server

- image_maker:
  - S3のPUTイベントで発火し、対象のコンテンツのサムネイル画像を生成するbatch

- notification:
  - SQSイベントで発火し、購入処理を行うbatch（現時点ではまだイベントを発火させるだけで具体的なロジックは未実装）

- rank:
  - EventBridgeイベントで発火するbatch

- reward:
  - EventBridgeイベントで発火するbatch
