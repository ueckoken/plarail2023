# k8sクラスタのセットアップ方法

1. k8sクラスタにログインできるようにする
2. Fluxのインストール
[https://](https://fluxcd.io/flux/installation/)
3. GithubのPAT (Personal Access Token)を取得する。
4. 以下のコマンドを実行する。

```zsh
flux bootstrap github \
--components-extra=image-reflector-controller,image-automation-controller \
--owner=ueckoken \
--repository=plarail2023 \
--branch=main \
--read-write-key \
--path=./infra/manifests
```
