function getData() {
    
    let name = document.getElementById('input-name').value
    let email = document.getElementById('input-email').value
    let phone = document.getElementById('input-phone').value
    let subject = document.getElementById('input-subject').value
    let message = document.getElementById('input-message').value
    
    
    if (name == '') {
        return alert('nama harus diisi bro...')
    }else if (email == '') {
        return alert('email harus diisi bro...')
    }else if (phone == '') {
        return alert('phone harus diisi bro...')
    }else if (subject == '') {
        return alert('subject harus diisi bro...')
    }else if (message == '') {
        return alert('message harus diisi bro...')
    }
    

    // console.log(name);
    // console.log(email);
    // console.log(phone);
    // console.log(subject);
    // console.log(message);

    let Mail = '@arifaicreative@gmail.com'
    let a = document.createElement('a')

    a.href = `mailto:${Mail}?subject=${subject}&body=Hi, my name ${name}. ${message}, this my phone number ${phone} and email: ${email}`
    a.click();

}