# !/bin/bash

read -p "ファイル名を入力してください: " name
read -p "struct名を入力してください: " structName

# もしプロジェクト名がなかったら終了
if [[ $name = ""  ]]; then
  echo "プロジェクト名を指定してください"
  exit;
fi

copy ./scripts/script-template/template_handler.go ./script-template/${name}.go


echo '🎉作成が完了しました🎉'
echo 'config/routerにroutingを記述しましょう。'
