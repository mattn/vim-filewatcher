let s:cmd = expand('<sfile>:h:h:gs!\\!/!') . '/filewatcher/filewatcher' . (has('win32') ? '.exe' : '')
if !filereadable(s:cmd)
  finish
endif

let s:template = {}

function! s:template.stop()
  call job_stop(self.job)
endfunction

function! filewatcher#watch(dir, cb)
  let ctx = copy(s:template)
  let ctx['dir'] = a:dir
  let ctx['job'] = job_start([s:cmd, a:dir], { 'out_cb': a:cb, 'out_mode': 'nl' })
  return ctx
endfunction
