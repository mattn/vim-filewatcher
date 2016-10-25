let s:cmd = expand('<sfile>:h:h:gs!\\!/!') . '/filewatcher/filewatcher.exe'
if !filereadable(s:cmd)
  finish
endif

function! filewatcher#watch(dir, cb)
  if exists('s:job')
    call job_stop(s:job)
  endif
  let s:job = job_start([s:cmd, a:dir], { 'out_cb': a:cb, 'out_mode': 'nl' })
endfunction
