podTemplate(label: "kubeless-go-sample", containers: [
    containerTemplate(name: 'kubeless', image: 'registry.gitlab.com/mvenezia/kubeless-cli:v1.0.0-alpha.7', ttyEnabled: true, command: 'cat')
  ]) {

    node("kubeless-go-sample") {
        stage('Get kubeconfig secret') {
            checkout scm

            container('kubeless') {
                def secrets = [
                    [
                        $class: 'VaultSecret', path: 'jobs/kubeless-go-sample', secretValues: [
                            [
                                $class: 'VaultSecretValue', envVar: 'KUBECONFIG', vaultKey: 'kubeconfig'
                            ]
                        ]
                    ]
                ]

                wrap([$class: 'VaultBuildWrapper', configuration: configuration, vaultSecrets: secrets]) {
                    sh 'echo $KUBECONFIG'
                }
            }
        }
    }
}