# GoAPI ハンズオン 最終更新日:1/28

Vue.js で実装したアプリのバックエンド部分です。<br>
<https://github.com/KEMPER0530/vue-handson>
![Architecture Diagram](https://d3p7ybpyhwl4ka.cloudfront.net/static/img/aws-Diagram.svg)

## 開発環境
- go 1.13.4 darwin/amd64
- Mysql:5.7
- Dockerにて環境構築
- Amazon Simple Email Serviceを利用したメール配信
- Firebaseを利用してAPIでJWT検証の実施
- メールリンク認証の機能を実装

## 機能一覧
- ログイン情報の取得API
- WORKの取得API
- CORS対応済
- クレジットカード情報登録API
- AWS(SES)を利用したメール送信
- アカウント登録API
- リバースプロキシ用にnginxを投入
- Firebaseを利用し認証の実施
- メールリンク認証

## 本番環境
- AWS(EC2)
- AWS(SES)
- AWS(S３)
- AWS(RDS)
- AWS(Route53)
- AWS(CloudFront)
- AWS(ELB)
- Firebase

## 今後実装したいこと
- CIの導入、テスト、デプロイの自動化
