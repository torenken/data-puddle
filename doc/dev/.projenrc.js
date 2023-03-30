const { awscdk } = require('projen');

const appName = 'data-puddle';

const project = new awscdk.AwsCdkTypeScriptApp({
  cdkVersion: '2.72.0',
  defaultReleaseBranch: 'main',
  authorName: 'Thomas Renken',
  repository: `git@github.com:torenken/${appName}.git`,
  name: appName,

  // deps: ['@aws-cdk/aws-lambda-go-alpha'],  /* Runtime dependencies of this module. */
  // devDeps: [],                             /* Build dependencies for this module. */

  context: {
    '@aws-cdk/core:newStyleStackSynthesis': true,
  },
  appEntrypoint: `${appName}-app.ts`,

  jestOptions: {
    jestVersion: '28',
  },

  typescriptVersion: '^4',

  //no-github-workflow
  github: false,

  license: 'Apache-2.0',
});

project.setScript('postinstall', 'touch node_modules/go.mod'); //This step is necessary so that go mod tidy ignores the cdk-go deps.
project.setScript('audit:check', 'yarn audit:level-high && yarn audit:level-critical');
project.setScript('audit:level-high', '/bin/bash -c \'yarn audit --groups dependencies; [[ $? -ge 8 ]] && exit 1 || exit 0\'');
project.setScript('audit:level-critical', '/bin/bash -c \'yarn audit --groups devDependencies; [[ $? -ge 16 ]] && exit 1 || exit 0\'');

project.synth();