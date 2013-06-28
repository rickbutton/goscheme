(begin 
(define
  rep
  (lambda () (begin
               (display "->")
               (let ((e (eval (read))))
                 (begin
                   (display "=>")
                   (print e))
               ))))
(for 1 (rep))
)
