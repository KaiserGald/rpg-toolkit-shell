module.exports = function(grunt) {
  //load npm tasks
  grunt.loadNpmTasks('grunt-bump');

  // grunt-bump config
  grunt.initConfig({
    bump: {
      options: {
        files: ['package.json'],
        updateConfigs: [],
        commit: true,
        commitMessage: 'Release v%VERSION%',
        commitFiles: ['-a'],
        createTag: true,
        tagName: 'v%VERSION%',
        tagMessage: 'Version %VERSION%',
        push: true,
        pushTo :'origin',
        gotDescribeOptions: '--tags --always --abbrev=1 --dirty=-d',
        globalReplace: false,
        prereleaseName: 'dev',
        metadata: '',
        regExp: false
      }
    },
  })
}
