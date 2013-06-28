(begin 
(define
  rep
  (lambda () (begin
               (display "->")
               (print (eval (read))))))
(for 1 (rep))
)
