import * as htmx from 'htmx.org'

import Alpine from 'alpinejs'
import Swal from 'sweetalert2'
// Add Alpine instance to window object.
window.Alpine = Alpine
window.Swal = Swal
window.htmx = htmx

// Start Alpine.
Alpine.start()

console.log('kauaka')
